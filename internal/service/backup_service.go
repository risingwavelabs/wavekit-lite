package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/risingwavelabs/wavekit/internal/apigen"
	"github.com/risingwavelabs/wavekit/internal/model"
	"github.com/risingwavelabs/wavekit/internal/model/querier"
	"github.com/risingwavelabs/wavekit/internal/modelctx"
	"github.com/risingwavelabs/wavekit/internal/utils"
)

const defaultBackupTaskTimeout = "30m"

func (s *Service) CreateClusterSnapshot(ctx context.Context, id int32, name string, orgID int32) (*apigen.Snapshot, error) {
	conn, err := s.getRisectlConn(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get risectl connection")
	}

	snapshotID, err := conn.MetaBackup(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create snapshot")
	}

	if err := s.m.CreateClusterSnapshot(ctx, querier.CreateClusterSnapshotParams{
		ClusterID:  id,
		SnapshotID: snapshotID,
		Name:       name,
	}); err != nil {
		return nil, errors.Wrapf(err, "failed to create snapshot")
	}

	return &apigen.Snapshot{
		ID:        snapshotID,
		Name:      name,
		ClusterID: id,
		CreatedAt: time.Now(),
	}, nil
}

func (s *Service) ListClusterSnapshots(ctx context.Context, id int32, orgID int32) ([]apigen.Snapshot, error) {
	snapshots, err := s.m.ListClusterSnapshots(ctx, id)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list cluster snapshots")
	}

	result := make([]apigen.Snapshot, len(snapshots))
	for i, snapshot := range snapshots {
		result[i] = apigen.Snapshot{
			ID:        snapshot.SnapshotID,
			Name:      snapshot.Name,
			ClusterID: snapshot.ClusterID,
			CreatedAt: snapshot.CreatedAt,
		}
	}

	return result, nil
}

func (s *Service) DeleteClusterSnapshot(ctx context.Context, id int32, snapshotID int64, orgID int32) error {
	conn, err := s.getRisectlConn(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "failed to get risectl connection")
	}

	if err := conn.DeleteSnapshot(ctx, snapshotID); err != nil {
		return errors.Wrapf(err, "failed to delete snapshot")
	}

	return nil
}

func (s *Service) UpdateClusterAutoBackupConfig(ctx context.Context, id int32, params apigen.AutoBackupConfig, orgID int32) error {
	cluster, err := s.m.GetOrgCluster(ctx, querier.GetOrgClusterParams{
		ID:             id,
		OrganizationID: orgID,
	})
	if err != nil {
		return errors.Wrapf(err, "failed to get cluster")
	}

	taskSpec := apigen.TaskSpec{
		Type: apigen.AutoBackup,
		AutoBackup: &apigen.TaskSpecAutoBackup{
			ClusterID:         cluster.ID,
			RetentionDuration: params.RetentionDuration,
		},
	}

	c, err := s.m.GetAutoBackupConfig(ctx, cluster.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// No existing auto backup config, create a new one and a new cron job
			if err := s.m.RunTransaction(ctx, func(txm model.ModelInterface) error {
				c := modelctx.NewModelctx(ctx, txm)
				taskID, err := s.taskstore.CreateCronJob(c, utils.Ptr(defaultBackupTaskTimeout), &orgID, params.CronExpression, taskSpec)
				if err != nil {
					return errors.Wrapf(err, "failed to create cron job")
				}
				if err := txm.CreateAutoBackupConfig(ctx, querier.CreateAutoBackupConfigParams{
					ClusterID: cluster.ID,
					TaskID:    taskID,
					Enabled:   true,
				}); err != nil {
					return errors.Wrapf(err, "failed to create auto backup config")
				}
				return nil
			}); err != nil {
				return errors.Wrapf(err, "failed to create new cluster auto backup config")
			}
			return nil
		}
		return errors.Wrapf(err, "failed to get auto backup config")
	}

	if err := s.m.RunTransaction(ctx, func(txm model.ModelInterface) error {
		mc := modelctx.NewModelctx(ctx, txm)
		if !params.Enabled {
			if err := s.taskstore.PauseCronJob(mc, c.TaskID); err != nil {
				return errors.Wrapf(err, "failed to pause cron job")
			}
		} else {
			if err := s.taskstore.ResumeCronJob(mc, c.TaskID); err != nil {
				return errors.Wrapf(err, "failed to resume cron job")
			}
		}
		if err := txm.UpdateAutoBackupConfig(ctx, querier.UpdateAutoBackupConfigParams{
			ClusterID: cluster.ID,
			Enabled:   params.Enabled,
		}); err != nil {
			return errors.Wrapf(err, "failed to update auto backup config")
		}
		if err := s.taskstore.UpdateCronJob(mc, c.TaskID, utils.Ptr(defaultBackupTaskTimeout), &orgID, params.CronExpression, taskSpec); err != nil {
			return errors.Wrapf(err, "failed to update cron job")
		}
		return nil
	}); err != nil {
		return errors.Wrapf(err, "failed to update cluster auto backup config")
	}
	return nil
}

func (s *Service) GetClusterAutoBackupConfig(ctx context.Context, id int32, orgID int32) (*apigen.AutoBackupConfig, error) {
	c, err := s.m.GetAutoBackupConfig(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &apigen.AutoBackupConfig{
				Enabled: false,
			}, nil
		}
		return nil, errors.Wrapf(err, "failed to get auto backup config")
	}
	task, err := s.m.GetTaskByID(ctx, c.TaskID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get task")
	}

	return &apigen.AutoBackupConfig{
		Enabled:           c.Enabled,
		CronExpression:    task.Attributes.Cronjob.CronExpression,
		RetentionDuration: task.Spec.AutoBackup.RetentionDuration,
	}, nil
}
