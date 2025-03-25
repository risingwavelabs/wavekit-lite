package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/risingwavelabs/wavekit/internal/apigen"
	"github.com/risingwavelabs/wavekit/internal/model/querier"
)

var (
	ErrMetricsStoreNotFound = fmt.Errorf("metrics store not found")
	ErrMetricsStoreInUse    = fmt.Errorf("metrics store is in use")
)

func metricsStoreToAPI(ms *querier.MetricsStore) *apigen.MetricsStore {
	return &apigen.MetricsStore{
		ID:        ms.ID,
		Name:      ms.Name,
		Spec:      &ms.Spec,
		CreatedAt: ms.CreatedAt.Time,
	}
}

// CreateMetricsStore creates a new metrics store
func (s *Service) CreateMetricsStore(ctx context.Context, req apigen.MetricsStoreCreate, organizationID int32) (*apigen.MetricsStore, error) {
	params := querier.CreateMetricsStoreParams{
		Name:           req.Name,
		Spec:           req.Spec,
		OrganizationID: organizationID,
	}

	ms, err := s.m.CreateMetricsStore(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create metrics store: %w", err)
	}

	return metricsStoreToAPI(ms), nil
}

func (s *Service) ListClustersByMetricsStoreID(ctx context.Context, id int32) ([]*apigen.Cluster, error) {
	clusters, err := s.m.ListClustersByMetricsStoreID(ctx, &id)
	if err != nil {
		return nil, fmt.Errorf("failed to list clusters by metrics store: %w", err)
	}

	apiClusters := make([]*apigen.Cluster, len(clusters))
	for i, cluster := range clusters {
		apiClusters[i] = clusterToApi(cluster)
	}
	return apiClusters, nil
}

func (s *Service) DeleteMetricsStore(ctx context.Context, id int32, organizationID int32) error {
	if err := s.m.DeleteMetricsStore(ctx, querier.DeleteMetricsStoreParams{
		ID:             id,
		OrganizationID: organizationID,
	}); err != nil {
		if err == sql.ErrNoRows {
			return ErrMetricsStoreNotFound
		}
		return fmt.Errorf("failed to delete metrics store: %w", err)
	}
	return nil
}

func (s *Service) GetMetricsStore(ctx context.Context, id int32, organizationID int32) (*apigen.MetricsStore, error) {
	ms, err := s.m.GetMetricsStoreByIDAndOrgID(ctx, querier.GetMetricsStoreByIDAndOrgIDParams{
		ID:             id,
		OrganizationID: organizationID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrMetricsStoreNotFound
		}
		return nil, fmt.Errorf("failed to get metrics store: %w", err)
	}

	return metricsStoreToAPI(ms), nil
}

func (s *Service) ListMetricsStores(ctx context.Context, organizationID int32) ([]*apigen.MetricsStore, error) {
	msList, err := s.m.ListMetricsStoresByOrgID(ctx, organizationID)
	if err != nil {
		return nil, fmt.Errorf("failed to list metrics stores: %w", err)
	}

	apiMsList := make([]*apigen.MetricsStore, len(msList))
	for i, ms := range msList {
		apiMsList[i] = metricsStoreToAPI(ms)
	}

	return apiMsList, nil
}

func (s *Service) UpdateMetricsStore(ctx context.Context, id int32, req apigen.MetricsStoreCreate, organizationID int32) (*apigen.MetricsStore, error) {
	_, err := s.m.GetMetricsStoreByIDAndOrgID(ctx, querier.GetMetricsStoreByIDAndOrgIDParams{
		ID:             id,
		OrganizationID: organizationID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrMetricsStoreNotFound
		}
		return nil, fmt.Errorf("failed to get metrics store: %w", err)
	}

	params := querier.UpdateMetricsStoreParams{
		ID:             id,
		Name:           req.Name,
		Spec:           req.Spec,
		OrganizationID: organizationID,
	}

	ms, err := s.m.UpdateMetricsStore(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update metrics store: %w", err)
	}

	return metricsStoreToAPI(ms), nil
}
