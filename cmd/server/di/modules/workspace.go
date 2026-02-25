// Package modules ...
package modules

import (
	"github.com/sudarsh1010/adfinis/internal/application/workspace"
	"github.com/sudarsh1010/adfinis/internal/infrastructure/persistence/bun/repositories"
	"go.uber.org/fx"
)

var WorkspaceModule = fx.Module("workspace",
	fx.Provide(repositories.NewWorkspaceRepository),
	fx.Provide(workspace.NewWorkspaceService),
)
