-- Migration: 001_init_schema
-- Description: Initial schema for vector DB management models
-- Created: 2026-02-28

-- Enable foreign keys (SQLite)
PRAGMA foreign_keys = ON;

-- Workspaces table
CREATE TABLE IF NOT EXISTS workspaces (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    is_default INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL
);

-- Provider Accounts table
CREATE TABLE IF NOT EXISTS provider_accounts (
    id TEXT PRIMARY KEY,
    workspace_id TEXT NOT NULL,
    provider TEXT NOT NULL,
    name TEXT NOT NULL,
    api_key_encrypted TEXT NOT NULL,
    extra_config_json TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_provider_accounts_workspace_id ON provider_accounts(workspace_id);

-- Connections table
CREATE TABLE IF NOT EXISTS connections (
    id TEXT PRIMARY KEY,
    workspace_id TEXT NOT NULL,
    name TEXT NOT NULL,
    provider TEXT NOT NULL,
    endpoint TEXT NOT NULL,
    api_key_encrypted TEXT,
    region TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL,
    last_connected_at DATETIME,
    FOREIGN KEY (workspace_id) REFERENCES workspaces(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_connections_workspace_id ON connections(workspace_id);

-- Database Namespaces table
CREATE TABLE IF NOT EXISTS database_namespaces (
    id TEXT PRIMARY KEY,
    connection_id TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (connection_id) REFERENCES connections(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_database_namespaces_connection_id ON database_namespaces(connection_id);

-- Embedding Profiles table
CREATE TABLE IF NOT EXISTS embedding_profiles (
    id TEXT PRIMARY KEY,
    provider_account_id TEXT NOT NULL,
    model_name TEXT NOT NULL,
    dimension INTEGER,
    metric TEXT NOT NULL,
    normalize INTEGER NOT NULL DEFAULT 0,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL,
    FOREIGN KEY (provider_account_id) REFERENCES provider_accounts(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_embedding_profiles_provider_account_id ON embedding_profiles(provider_account_id);

-- Collections table
CREATE TABLE IF NOT EXISTS collections (
    id TEXT PRIMARY KEY,
    database_namespace_id TEXT NOT NULL,
    name TEXT NOT NULL,
    embedding_profile_id TEXT NOT NULL,
    vector_count INTEGER NOT NULL DEFAULT 0,
    metadata_schema_json TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL,
    last_synced_at DATETIME,
    FOREIGN KEY (database_namespace_id) REFERENCES database_namespaces(id) ON DELETE CASCADE,
    FOREIGN KEY (embedding_profile_id) REFERENCES embedding_profiles(id) ON DELETE CASCADE,
    UNIQUE (database_namespace_id, name)
);

-- Cross-tenant integrity trigger for collections
-- Ensures embedding profile belongs to same workspace as namespace
CREATE TRIGGER IF NOT EXISTS check_collection_workspace_match
BEFORE INSERT ON collections
FOR EACH ROW
BEGIN
    SELECT RAISE(ABORT, 'Cross-tenant violation: embedding profile must be in same workspace as namespace')
    WHERE (
        SELECT dn.connection_id FROM database_namespaces dn WHERE dn.id = NEW.database_namespace_id
    ) NOT IN (
        SELECT c.id FROM connections c WHERE c.workspace_id = (
            SELECT pa.workspace_id FROM embedding_profiles ep
            JOIN provider_accounts pa ON ep.provider_account_id = pa.id
            WHERE ep.id = NEW.embedding_profile_id
        )
    );
END;

-- Also check on update
CREATE TRIGGER IF NOT EXISTS check_collection_workspace_match_update
BEFORE UPDATE ON collections
FOR EACH ROW
WHEN NEW.database_namespace_id != OLD.database_namespace_id OR NEW.embedding_profile_id != OLD.embedding_profile_id
BEGIN
    SELECT RAISE(ABORT, 'Cross-tenant violation: embedding profile must be in same workspace as namespace')
    WHERE (
        SELECT dn.connection_id FROM database_namespaces dn WHERE dn.id = NEW.database_namespace_id
    ) NOT IN (
        SELECT c.id FROM connections c WHERE c.workspace_id = (
            SELECT pa.workspace_id FROM embedding_profiles ep
            JOIN provider_accounts pa ON ep.provider_account_id = pa.id
            WHERE ep.id = NEW.embedding_profile_id
        )
    );
END;

CREATE INDEX IF NOT EXISTS idx_collections_database_namespace_id ON collections(database_namespace_id);
CREATE INDEX IF NOT EXISTS idx_collections_embedding_profile_id ON collections(embedding_profile_id);
