package modules

import (
	"github.com/sudarsh1010/adfinis/internal/application/collection"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/persistence/bun/repositories"
	"go.uber.org/fx"
)

var CollectionModule = fx.Module("collection",
	fx.Provide(
		fx.Annotate(
			repositories.NewCollectionRepository,
			fx.As(new(collection.Repository)),
		),
	),
	fx.Provide(collection.NewService),
)
