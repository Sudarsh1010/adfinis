// Package web provides embedded static files for the frontend.
package web

import (
	"embed"
	"io/fs"
)

//go:embed build
var staticFiles embed.FS

// Static returns the embedded static files with the build prefix stripped.
func Static() (fs.FS, error) {
	return fs.Sub(staticFiles, "build")
}
