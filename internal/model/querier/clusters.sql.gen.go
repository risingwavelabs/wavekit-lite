// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: clusters.sql

package querier

import (
	"context"
)

const createCluster = `-- name: CreateCluster :one
INSERT INTO clusters (
    organization_id,
    name,
    host,
    sql_port,
    meta_port
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, organization_id, name, host, sql_port, meta_port, created_at, updated_at
`

type CreateClusterParams struct {
	OrganizationID int32
	Name           string
	Host           string
	SqlPort        int32
	MetaPort       int32
}

func (q *Queries) CreateCluster(ctx context.Context, arg CreateClusterParams) (*Cluster, error) {
	row := q.db.QueryRow(ctx, createCluster,
		arg.OrganizationID,
		arg.Name,
		arg.Host,
		arg.SqlPort,
		arg.MetaPort,
	)
	var i Cluster
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Name,
		&i.Host,
		&i.SqlPort,
		&i.MetaPort,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const deleteOrgCluster = `-- name: DeleteOrgCluster :exec
DELETE FROM clusters
WHERE id = $1 AND organization_id = $2
`

type DeleteOrgClusterParams struct {
	ID             int32
	OrganizationID int32
}

func (q *Queries) DeleteOrgCluster(ctx context.Context, arg DeleteOrgClusterParams) error {
	_, err := q.db.Exec(ctx, deleteOrgCluster, arg.ID, arg.OrganizationID)
	return err
}

const getClusterByID = `-- name: GetClusterByID :one
SELECT id, organization_id, name, host, sql_port, meta_port, created_at, updated_at FROM clusters
WHERE id = $1
`

func (q *Queries) GetClusterByID(ctx context.Context, id int32) (*Cluster, error) {
	row := q.db.QueryRow(ctx, getClusterByID, id)
	var i Cluster
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Name,
		&i.Host,
		&i.SqlPort,
		&i.MetaPort,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getOrgCluster = `-- name: GetOrgCluster :one
SELECT id, organization_id, name, host, sql_port, meta_port, created_at, updated_at FROM clusters
WHERE id = $1 AND organization_id = $2
`

type GetOrgClusterParams struct {
	ID             int32
	OrganizationID int32
}

func (q *Queries) GetOrgCluster(ctx context.Context, arg GetOrgClusterParams) (*Cluster, error) {
	row := q.db.QueryRow(ctx, getOrgCluster, arg.ID, arg.OrganizationID)
	var i Cluster
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Name,
		&i.Host,
		&i.SqlPort,
		&i.MetaPort,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const initCluster = `-- name: InitCluster :one
INSERT INTO clusters (
    organization_id,
    name,
    host,
    sql_port,
    meta_port
) VALUES (
    $1, $2, $3, $4, $5
) ON CONFLICT (organization_id, name) DO UPDATE SET updated_at = CURRENT_TIMESTAMP
RETURNING id, organization_id, name, host, sql_port, meta_port, created_at, updated_at
`

type InitClusterParams struct {
	OrganizationID int32
	Name           string
	Host           string
	SqlPort        int32
	MetaPort       int32
}

func (q *Queries) InitCluster(ctx context.Context, arg InitClusterParams) (*Cluster, error) {
	row := q.db.QueryRow(ctx, initCluster,
		arg.OrganizationID,
		arg.Name,
		arg.Host,
		arg.SqlPort,
		arg.MetaPort,
	)
	var i Cluster
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Name,
		&i.Host,
		&i.SqlPort,
		&i.MetaPort,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const listOrgClusters = `-- name: ListOrgClusters :many
SELECT id, organization_id, name, host, sql_port, meta_port, created_at, updated_at FROM clusters
WHERE organization_id = $1
ORDER BY name
`

func (q *Queries) ListOrgClusters(ctx context.Context, organizationID int32) ([]*Cluster, error) {
	rows, err := q.db.Query(ctx, listOrgClusters, organizationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Cluster
	for rows.Next() {
		var i Cluster
		if err := rows.Scan(
			&i.ID,
			&i.OrganizationID,
			&i.Name,
			&i.Host,
			&i.SqlPort,
			&i.MetaPort,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrgCluster = `-- name: UpdateOrgCluster :one
UPDATE clusters
SET
    name = $3,
    host = $4,
    sql_port = $5,
    meta_port = $6,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND organization_id = $2
RETURNING id, organization_id, name, host, sql_port, meta_port, created_at, updated_at
`

type UpdateOrgClusterParams struct {
	ID             int32
	OrganizationID int32
	Name           string
	Host           string
	SqlPort        int32
	MetaPort       int32
}

func (q *Queries) UpdateOrgCluster(ctx context.Context, arg UpdateOrgClusterParams) (*Cluster, error) {
	row := q.db.QueryRow(ctx, updateOrgCluster,
		arg.ID,
		arg.OrganizationID,
		arg.Name,
		arg.Host,
		arg.SqlPort,
		arg.MetaPort,
	)
	var i Cluster
	err := row.Scan(
		&i.ID,
		&i.OrganizationID,
		&i.Name,
		&i.Host,
		&i.SqlPort,
		&i.MetaPort,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
