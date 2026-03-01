-- Migration: 002_add_encryption_keys
-- Description: Table to store the encryption key for nacl/secretbox encryption
-- Created: 2026-03-01

-- Table to store the auto-generated encryption key
-- Only one row should exist (id = "default")
CREATE TABLE IF NOT EXISTS encryption_keys (
    id TEXT PRIMARY KEY,
    key_data TEXT NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
