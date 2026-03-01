package modules

import (
	"github.com/sudarsh1010/adfinis/internal/infrastructure/milvus"
	"go.uber.org/fx"
)

var MilvusModule = fx.Module("milvus",
	fx.Provide(milvus.NewAdapter),
)
