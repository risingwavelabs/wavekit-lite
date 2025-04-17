// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: opaque_keys.sql

package querier

import (
	"context"
)

const createOpaqueKey = `-- name: CreateOpaqueKey :one
INSERT INTO opaque_keys (user_id, key) VALUES ($1, $2) RETURNING id
`

type CreateOpaqueKeyParams struct {
	UserID int32
	Key    []byte
}

func (q *Queries) CreateOpaqueKey(ctx context.Context, arg CreateOpaqueKeyParams) (int64, error) {
	row := q.db.QueryRow(ctx, createOpaqueKey, arg.UserID, arg.Key)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteOpaqueKey = `-- name: DeleteOpaqueKey :exec
DELETE FROM opaque_keys WHERE id = $1
`

func (q *Queries) DeleteOpaqueKey(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteOpaqueKey, id)
	return err
}

const deleteOpaqueKeys = `-- name: DeleteOpaqueKeys :exec
DELETE FROM opaque_keys WHERE user_id = $1
`

func (q *Queries) DeleteOpaqueKeys(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteOpaqueKeys, userID)
	return err
}

const getOpaqueKey = `-- name: GetOpaqueKey :one
SELECT key FROM opaque_keys WHERE id = $1
`

func (q *Queries) GetOpaqueKey(ctx context.Context, id int64) ([]byte, error) {
	row := q.db.QueryRow(ctx, getOpaqueKey, id)
	var key []byte
	err := row.Scan(&key)
	return key, err
}
