package di

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type Database struct {
	*bun.DB
}

func NewDatabase(
	ctx context.Context,
	cfg *Config,
	logger *Logger,
) (*Database, error) {
	logger.Info("Initializing database", "path", cfg.DBPath)

	var err error

	sqldb, err := sql.Open(sqliteshim.ShimName, cfg.DBPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite: %w", err)
	}

	// Connection pooling
	sqldb.SetMaxOpenConns(BunMaxOpenConns)
	sqldb.SetMaxIdleConns(BunMaxIdleConns)
	sqldb.SetConnMaxLifetime(BunConnMaxLifetime)
	sqldb.SetConnMaxIdleTime(BunConnMaxIdleTime)

	// Create Bun DB with debug hook in development
	// var db *bun.DB
	// if cfg.Env == "development" {
	// 	db = bun.NewDB(sqldb, &bundebug.QueryHook{})
	// } else {
	db := bun.NewDB(sqldb, sqlitedialect.New())
	// }

	// Test connection
	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	logger.Info("Database initialized successfully")
	return &Database{db}, nil
}

func (d *Database) Close(_ context.Context) error {
	return d.DB.Close()
}
