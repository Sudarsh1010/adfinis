package modules

import (
	"github.com/sudarsh1010/adfinis/internal/application/databasenamespace"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/persistence/bun/repositories"
	"go.uber.org/fx"
)

var DatabaseNamespaceModule = fx.Module("databasenamespace",
	fx.Provide(
		fx.Annotate(
			repositories.NewDatabaseNamespaceRepository,
			fx.As(new(databasenamespace.Repository)),
		),
	),
	fx.Provide(databasenamespace.NewDatabaseNamespaceService),
)
