package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/risingwavelabs/wavekit/internal/apigen"
	"github.com/risingwavelabs/wavekit/internal/model/querier"
	"github.com/robfig/cron/v3"
)

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

	parser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	cron, err := parser.Parse(params.CronExpression)
	if err != nil {
		return errors.Wrapf(err, "failed to parse cron expression")
	}
	nextTime := cron.Next(time.Now())

	taskAttributes := apigen.TaskAttributes{
		OrgID: &orgID,
		Cronjob: &apigen.TaskCronjob{
			CronExpression: params.CronExpression,
		},
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
			_, err := s.m.CreateTask(ctx, querier.CreateTaskParams{
				Attributes: taskAttributes,
				Spec:       taskSpec,
				StartedAt:  &nextTime,
			})
			if err != nil {
				return errors.Wrapf(err, "failed to create task")
			}
			return nil
		}
		return errors.Wrapf(err, "failed to get auto backup config")
	}

	if err := s.m.UpdateTask(ctx, querier.UpdateTaskParams{
		ID:         c.TaskID,
		Attributes: taskAttributes,
		Spec:       taskSpec,
		StartedAt:  &nextTime,
	}); err != nil {
		return errors.Wrapf(err, "failed to update task")
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
