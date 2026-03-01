package modules

import (
	"github.com/sudarsh1010/adfinis/internal/application/connection"
	"github.com/sudarsh1010/adfinis/internal/domain/milvus"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/crypto"
	milvusadapter "github.com/sudarsh1010/adfinis/internal/infrastructure/milvus"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/persistence/bun/repositories"
	"go.uber.org/fx"
)

var ConnectionModule = fx.Module("connection",
	// Repository with interface annotation
	fx.Provide(
		fx.Annotate(
			repositories.NewConnectionRepository,
			fx.As(new(connection.Repository)),
		),
	),
	// Crypto service
	fx.Provide(crypto.NewService),
	// Milvus adapter with interface annotation
	fx.Provide(
		fx.Annotate(
			milvusadapter.NewAdapter,
			fx.As(new(milvus.MilvusClient)),
		),
	),
	// Connection service
	fx.Provide(connection.NewService),
)
