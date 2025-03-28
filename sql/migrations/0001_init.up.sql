BEGIN;

CREATE TABLE IF NOT EXISTS organizations (
    id            SERIAL,
    name          TEXT        NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at    TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    UNIQUE (name),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id              SERIAL,
    name            TEXT        NOT NULL,
    password_hash   TEXT        NOT NULL,
    password_salt   TEXT        NOT NULL,
    organization_id INTEGER     NOT NULL REFERENCES organizations(id),
    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    UNIQUE (name),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS organization_owners (
    user_id         INTEGER     NOT NULL REFERENCES users(id),
    organization_id INTEGER     NOT NULL REFERENCES organizations(id),
    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (user_id, organization_id)
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
    id              SERIAL,
    user_id         INTEGER     NOT NULL REFERENCES users(id),
    token           TEXT        NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (id),
    UNIQUE (user_id, token)
);

CREATE TABLE IF NOT EXISTS clusters (
    id              SERIAL,
    organization_id INTEGER     NOT NULL REFERENCES organizations(id),

    name            TEXT        NOT NULL,
    host            TEXT        NOT NULL,
    sql_port        INTEGER     NOT NULL,
    meta_port       INTEGER     NOT NULL,
    http_port       INTEGER     NOT NULL,
    version         TEXT        NOT NULL,
    
    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    UNIQUE (organization_id, name),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS database_connections (
    id              SERIAL,
    organization_id INTEGER     NOT NULL REFERENCES organizations(id),

    name            TEXT        NOT NULL,
    cluster_id      INTEGER     NOT NULL REFERENCES clusters(id),
    username        TEXT        NOT NULL,
    password        TEXT,
    database        TEXT        NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    UNIQUE (organization_id, name),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS cluster_snapshots (
    cluster_id      INTEGER     NOT NULL REFERENCES clusters(id),
    snapshot_id     BIGINT      NOT NULL,
    name            TEXT        NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (cluster_id, snapshot_id)
);

CREATE TABLE IF NOT EXISTS cluster_diagnostics (
    id              SERIAL,
    cluster_id      INTEGER     NOT NULL REFERENCES clusters(id),
    content         TEXT        NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS auto_backup_configs (
    cluster_id      INTEGER     NOT NULL REFERENCES clusters(id),
    enabled         BOOLEAN     NOT NULL DEFAULT FALSE,
    cron_expression TEXT        NOT NULL,
    keep_last       INTEGER     NOT NULL DEFAULT 1,

    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (cluster_id)
);

CREATE TABLE IF NOT EXISTS auto_diagnostics_configs (
    cluster_id         INTEGER     NOT NULL REFERENCES clusters(id),
    enabled            BOOLEAN     NOT NULL DEFAULT FALSE,
    cron_expression    TEXT        NOT NULL,
    retention_duration TEXT, -- Null means no expiration

    created_at         TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at         TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (cluster_id)
);

CREATE TABLE IF NOT EXISTS org_settings (
    organization_id INTEGER     NOT NULL REFERENCES organizations(id),
    timezone        TEXT        NOT NULL DEFAULT 'UTC',

    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (organization_id)
);

COMMIT;
