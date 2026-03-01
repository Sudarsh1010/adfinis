// Package di ...
package di

import (
	"github.com/sudarsh1010/adfinis/cmd/server/di/modules"
	"github.com/sudarsh1010/adfinis/internal/interface/http/handlers"
	"github.com/uptrace/bun"
	"go.uber.org/fx"
)

// Module is the root FX module for the server.
var Module = fx.Options(
	// Core infrastructure
	fx.Provide(NewConfig),
	fx.Provide(NewLogger),
	fx.Provide(NewDatabase),

	// Extract *bun.DB from *Database for repositories
	fx.Provide(func(db *Database) *bun.DB { return db.DB }),

	// Migration
	fx.Invoke(RunMigrations),
	// fx.Invoke(RunMigrations),

	// HTTP layer
	fx.Provide(NewStaticFS),
	fx.Provide(NewSPAHandler),
	fx.Provide(NewRouter),
	fx.Provide(NewServer),

	// Handlers
	fx.Provide(handlers.NewHealthHandler),
	// fx.Provide(NewWorkspaceHandler),
	// fx.Provide(NewConnectionHandler),

	// Modules
	modules.WorkspaceModule,
	modules.ConnectionModule,
	modules.DatabaseNamespaceModule,
	modules.CollectionModule,

	// Handlers
	fx.Provide(handlers.NewWorkspaceHandler),
	fx.Provide(handlers.NewConnectionHandler),
	fx.Provide(handlers.NewDatabaseNamespaceHandler),
	fx.Provide(handlers.NewCollectionHandler),
)
