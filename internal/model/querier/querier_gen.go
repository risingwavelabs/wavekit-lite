// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querier

import (
	"context"
)

type Querier interface {
	CreateCluster(ctx context.Context, arg CreateClusterParams) (*Cluster, error)
	CreateDatabaseConnection(ctx context.Context, arg CreateDatabaseConnectionParams) (*DatabaseConnection, error)
	CreateOrganization(ctx context.Context, name string) (*Organization, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (*User, error)
	DeleteCluster(ctx context.Context, id int32) error
	DeleteDatabaseConnection(ctx context.Context, id int32) error
	DeleteOrganization(ctx context.Context, id int32) error
	DeleteRefreshToken(ctx context.Context, arg DeleteRefreshTokenParams) error
	DeleteUserByName(ctx context.Context, name string) error
	GetCluster(ctx context.Context, id int32) (*Cluster, error)
	GetDatabaseConnection(ctx context.Context, id int32) (*DatabaseConnection, error)
	GetOrganization(ctx context.Context, id int32) (*Organization, error)
	GetRefreshToken(ctx context.Context, arg GetRefreshTokenParams) (*RefreshToken, error)
	GetUser(ctx context.Context, id int32) (*User, error)
	GetUserByName(ctx context.Context, name string) (*User, error)
	ListClusters(ctx context.Context, organizationID int32) ([]*Cluster, error)
	ListDatabaseConnections(ctx context.Context, organizationID int32) ([]*DatabaseConnection, error)
	ListOrganizations(ctx context.Context) ([]*Organization, error)
	UpdateCluster(ctx context.Context, arg UpdateClusterParams) (*Cluster, error)
	UpdateDatabaseConnection(ctx context.Context, arg UpdateDatabaseConnectionParams) (*DatabaseConnection, error)
	UpdateOrganization(ctx context.Context, arg UpdateOrganizationParams) (*Organization, error)
	UpsertRefreshToken(ctx context.Context, arg UpsertRefreshTokenParams) error
}

var _ Querier = (*Queries)(nil)
