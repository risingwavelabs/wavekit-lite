// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querier

import (
	"context"
)

type Querier interface {
	CreateAutoBackupConfig(ctx context.Context, arg CreateAutoBackupConfigParams) error
	CreateAutoDiagnosticsConfig(ctx context.Context, arg CreateAutoDiagnosticsConfigParams) error
	CreateCluster(ctx context.Context, arg CreateClusterParams) (*Cluster, error)
	CreateClusterDiagnostic(ctx context.Context, arg CreateClusterDiagnosticParams) (*ClusterDiagnostic, error)
	CreateClusterSnapshot(ctx context.Context, arg CreateClusterSnapshotParams) error
	CreateDatabaseConnection(ctx context.Context, arg CreateDatabaseConnectionParams) (*DatabaseConnection, error)
	CreateMetricsStore(ctx context.Context, arg CreateMetricsStoreParams) (*MetricsStore, error)
	CreateOrgSettings(ctx context.Context, arg CreateOrgSettingsParams) error
	DeleteAllOrgDatabaseConnectionsByClusterID(ctx context.Context, arg DeleteAllOrgDatabaseConnectionsByClusterIDParams) error
	DeleteClusterDiagnostic(ctx context.Context, id int32) error
	DeleteClusterSnapshot(ctx context.Context, arg DeleteClusterSnapshotParams) error
	DeleteMetricsStore(ctx context.Context, arg DeleteMetricsStoreParams) error
	DeleteOrgCluster(ctx context.Context, arg DeleteOrgClusterParams) error
	DeleteOrgDatabaseConnection(ctx context.Context, arg DeleteOrgDatabaseConnectionParams) error
	GetAllOrgDatabseConnectionsByClusterID(ctx context.Context, arg GetAllOrgDatabseConnectionsByClusterIDParams) ([]*DatabaseConnection, error)
	GetAutoBackupConfig(ctx context.Context, clusterID int32) (*AutoBackupConfig, error)
	GetAutoDiagnosticsConfig(ctx context.Context, clusterID int32) (*AutoDiagnosticsConfig, error)
	GetClusterByID(ctx context.Context, id int32) (*Cluster, error)
	GetClusterDiagnostic(ctx context.Context, id int32) (*ClusterDiagnostic, error)
	GetDatabaseConnectionByID(ctx context.Context, id int32) (*DatabaseConnection, error)
	GetMetricsStore(ctx context.Context, id int32) (*MetricsStore, error)
	GetMetricsStoreByIDAndOrgID(ctx context.Context, arg GetMetricsStoreByIDAndOrgIDParams) (*MetricsStore, error)
	GetOrgCluster(ctx context.Context, arg GetOrgClusterParams) (*Cluster, error)
	GetOrgDatabaseByID(ctx context.Context, arg GetOrgDatabaseByIDParams) (*DatabaseConnection, error)
	GetOrgDatabaseConnection(ctx context.Context, arg GetOrgDatabaseConnectionParams) (*DatabaseConnection, error)
	GetOrgSettings(ctx context.Context, orgID int32) (*OrgSetting, error)
	InitCluster(ctx context.Context, arg InitClusterParams) (*Cluster, error)
	InitDatabaseConnection(ctx context.Context, arg InitDatabaseConnectionParams) (*DatabaseConnection, error)
	InitMetricsStore(ctx context.Context, arg InitMetricsStoreParams) (*MetricsStore, error)
	ListClusterDiagnostics(ctx context.Context, clusterID int32) ([]*ListClusterDiagnosticsRow, error)
	ListClusterSnapshots(ctx context.Context, clusterID int32) ([]*ClusterSnapshot, error)
	ListClustersByMetricsStoreID(ctx context.Context, metricsStoreID *int32) ([]*Cluster, error)
	ListMetricsStoresByOrgID(ctx context.Context, orgID int32) ([]*MetricsStore, error)
	ListOrgClusters(ctx context.Context, orgID int32) ([]*Cluster, error)
	ListOrgDatabaseConnections(ctx context.Context, orgID int32) ([]*DatabaseConnection, error)
	RemoveClusterMetricsStoreID(ctx context.Context, arg RemoveClusterMetricsStoreIDParams) error
	UpdateAutoBackupConfig(ctx context.Context, arg UpdateAutoBackupConfigParams) error
	UpdateAutoDiagnosticsConfig(ctx context.Context, arg UpdateAutoDiagnosticsConfigParams) error
	UpdateMetricsStore(ctx context.Context, arg UpdateMetricsStoreParams) (*MetricsStore, error)
	UpdateOrgCluster(ctx context.Context, arg UpdateOrgClusterParams) (*Cluster, error)
	UpdateOrgDatabaseConnection(ctx context.Context, arg UpdateOrgDatabaseConnectionParams) (*DatabaseConnection, error)
}

var _ Querier = (*Queries)(nil)
