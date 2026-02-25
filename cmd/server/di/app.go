// Package di ...
package di

import (
	"github.com/sudarsh1010/adfinis/internal/interface/http/handlers"
	"go.uber.org/fx"
)

// Module is the root FX module for the server.
var Module = fx.Options(
	// Core infrastructure
	fx.Provide(NewConfig),
	fx.Provide(NewLogger),
	fx.Provide(NewDatabase),

	// Migration
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
)
