package di

import (
	"io/fs"

	"github.com/sudarsh1010/adfinis/internal/interface/http/handlers"
	"github.com/sudarsh1010/adfinis/web"
	"go.uber.org/fx"
)

// StaticFSParams holds dependencies for creating the static FS.
type StaticFSParams struct {
	fx.In

	Config *Config
}

// NewStaticFS creates a [fs.FS] from the embedded static files.
// It uses the web package which embeds the build directory.
func NewStaticFS() (fs.FS, error) {
	return web.Static()
}

// NewSPAHandler creates a new SPA handler with the embedded static files.
func NewSPAHandler(
	p StaticFSParams,
	staticFS fs.FS,
) (*handlers.SPAHandler, error) {
	isDev := p.Config.Env == "development"
	return handlers.NewSPAHandler(staticFS, isDev)
}
