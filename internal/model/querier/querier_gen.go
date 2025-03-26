// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querier

import (
	"context"

	"github.com/risingwavelabs/wavekit/internal/apigen"
)

type Querier interface {
	CreateCluster(ctx context.Context, arg CreateClusterParams) (*Cluster, error)
	CreateClusterDiagnostic(ctx context.Context, arg CreateClusterDiagnosticParams) (*ClusterDiagnostic, error)
	CreateClusterSnapshot(ctx context.Context, arg CreateClusterSnapshotParams) error
	CreateDatabaseConnection(ctx context.Context, arg CreateDatabaseConnectionParams) (*DatabaseConnection, error)
	CreateMetricsStore(ctx context.Context, arg CreateMetricsStoreParams) (*MetricsStore, error)
	CreateOrganization(ctx context.Context, name string) (*Organization, error)
	CreateOrganizationOwner(ctx context.Context, arg CreateOrganizationOwnerParams) error
	CreateTask(ctx context.Context, arg CreateTaskParams) (*Task, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (*User, error)
	DeleteAllOrgDatabaseConnectionsByClusterID(ctx context.Context, arg DeleteAllOrgDatabaseConnectionsByClusterIDParams) error
	DeleteClusterSnapshot(ctx context.Context, arg DeleteClusterSnapshotParams) error
	DeleteMetricsStore(ctx context.Context, arg DeleteMetricsStoreParams) error
	DeleteOrgCluster(ctx context.Context, arg DeleteOrgClusterParams) error
	DeleteOrgDatabaseConnection(ctx context.Context, arg DeleteOrgDatabaseConnectionParams) error
	DeleteOrganization(ctx context.Context, id int32) error
	DeleteRefreshToken(ctx context.Context, arg DeleteRefreshTokenParams) error
	DeleteUserByName(ctx context.Context, name string) error
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
	GetOrganization(ctx context.Context, id int32) (*Organization, error)
	GetRefreshToken(ctx context.Context, arg GetRefreshTokenParams) (*RefreshToken, error)
	GetUser(ctx context.Context, id int32) (*User, error)
	GetUserByName(ctx context.Context, name string) (*User, error)
	InitCluster(ctx context.Context, arg InitClusterParams) (*Cluster, error)
	InitDatabaseConnection(ctx context.Context, arg InitDatabaseConnectionParams) (*DatabaseConnection, error)
	InsertEvent(ctx context.Context, spec apigen.EventSpec) (*Event, error)
	ListClusterDiagnostics(ctx context.Context, clusterID int32) ([]*ListClusterDiagnosticsRow, error)
	ListClusterSnapshots(ctx context.Context, clusterID int32) ([]*ClusterSnapshot, error)
	ListClustersByMetricsStoreID(ctx context.Context, metricsStoreID *int32) ([]*Cluster, error)
	ListMetricsStoresByOrgID(ctx context.Context, organizationID int32) ([]*MetricsStore, error)
	ListOrgClusters(ctx context.Context, organizationID int32) ([]*Cluster, error)
	ListOrgDatabaseConnections(ctx context.Context, organizationID int32) ([]*DatabaseConnection, error)
	ListOrganizations(ctx context.Context) ([]*Organization, error)
	LockTask(ctx context.Context, arg LockTaskParams) (*Task, error)
	PullTask(ctx context.Context, workerName *string) (*Task, error)
	RemoveClusterMetricsStoreID(ctx context.Context, arg RemoveClusterMetricsStoreIDParams) error
	SendWorkerHeartbeat(ctx context.Context, arg SendWorkerHeartbeatParams) error
	UpdateMetricsStore(ctx context.Context, arg UpdateMetricsStoreParams) (*MetricsStore, error)
	UpdateOrgCluster(ctx context.Context, arg UpdateOrgClusterParams) (*Cluster, error)
	UpdateOrgDatabaseConnection(ctx context.Context, arg UpdateOrgDatabaseConnectionParams) (*DatabaseConnection, error)
	UpdateOrganization(ctx context.Context, arg UpdateOrganizationParams) (*Organization, error)
	UpdateTaskMetadata(ctx context.Context, arg UpdateTaskMetadataParams) (*Task, error)
	UpdateTaskSpec(ctx context.Context, arg UpdateTaskSpecParams) (*Task, error)
	UpsertAutoBackupConfig(ctx context.Context, arg UpsertAutoBackupConfigParams) error
	UpsertAutoDiagnosticsConfig(ctx context.Context, arg UpsertAutoDiagnosticsConfigParams) error
	UpsertRefreshToken(ctx context.Context, arg UpsertRefreshTokenParams) error
}

var _ Querier = (*Queries)(nil)
