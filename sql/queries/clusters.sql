-- name: CreateCluster :one
INSERT INTO clusters (
    org_id,
    name,
    host,
    sql_port,
    meta_port,
    http_port,
    version, 
    metrics_store_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: InitCluster :one
INSERT INTO clusters (
    org_id,
    name,
    host,
    sql_port,
    meta_port,
    http_port,
    version,
    metrics_store_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) ON CONFLICT (org_id, name) DO UPDATE 
    SET 
        host = EXCLUDED.host,
        sql_port = EXCLUDED.sql_port,
        meta_port = EXCLUDED.meta_port,
        http_port = EXCLUDED.http_port,
        version = EXCLUDED.version,
        metrics_store_id = EXCLUDED.metrics_store_id,
        updated_at = CURRENT_TIMESTAMP
RETURNING *;

-- name: GetOrgCluster :one
SELECT * FROM clusters
WHERE id = $1 AND org_id = $2;

-- name: ListOrgClusters :many
SELECT * FROM clusters
WHERE org_id = $1
ORDER BY name;

-- name: UpdateOrgCluster :one
UPDATE clusters
SET
    name = $3,
    host = $4,
    sql_port = $5,
    meta_port = $6,
    http_port = $7,
    version = $8,
    metrics_store_id = $9,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND org_id = $2
RETURNING *;

-- name: DeleteOrgCluster :exec
DELETE FROM clusters
WHERE id = $1 AND org_id = $2;

-- name: GetClusterByID :one
SELECT * FROM clusters
WHERE id = $1;

-- name: ListClustersByMetricsStoreID :many
SELECT * FROM clusters
WHERE metrics_store_id = $1
ORDER BY name;


-- name: RemoveClusterMetricsStoreID :exec
UPDATE clusters
SET metrics_store_id = NULL
WHERE id = $1 AND org_id = $2;
