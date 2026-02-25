package handlers

import (
	"io"
	"io/fs"
	"mime"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strings"
	"time"
)

const (
	DialTimeout     = 10 * time.Second
	ProxyViteTimout = 30 * time.Second
	BidiProxySize   = 2
)

// SPAHandler serves a Single Page Application with:
// - Development mode: proxies to Vite dev server (localhost:5173) with WebSocket HMR support
// - Production mode: serves embedded static files with SPA fallback.
type SPAHandler struct {
	staticFS  fs.FS
	indexHTML []byte
	isDev     bool
	proxy     *httputil.ReverseProxy
}

// NewSPAHandler creates a new SPA handler.
// In development mode (isDev=true), it proxies to Vite dev server.
// In production mode, it serves files from the provided [fs.FS].
func NewSPAHandler(staticFS fs.FS, isDev bool) (*SPAHandler, error) {
	h := &SPAHandler{
		staticFS: staticFS,
		isDev:    isDev,
	}

	if isDev {
		target, err := url.Parse("http://localhost:5173")
		if err != nil {
			return nil, err
		}

		proxy := httputil.NewSingleHostReverseProxy(target)

		// Secure transport (prevents SSRF false positives)
		proxy.Transport = &http.Transport{
			Proxy: nil,
			DialContext: (&net.Dialer{
				Timeout: DialTimeout,
			}).DialContext,
		}

		// Optional: better error handling
		proxy.ErrorHandler = func(w http.ResponseWriter, _ *http.Request, _ error) {
			http.Error(
				w,
				"Vite dev server not reachable",
				http.StatusBadGateway,
			)
		}

		h.proxy = proxy
	} else {
		indexHTML, err := fs.ReadFile(staticFS, "index.html")
		if err != nil {
			return nil, err
		}
		h.indexHTML = indexHTML
	}

	return h, nil
}

// ServeHTTP handles HTTP requests for the SPA.
func (h *SPAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.isDev {
		h.proxy.ServeHTTP(
			w,
			r,
		) // #nosec G704 -- target host is hardcoded and transport locked to localhost
		return
	}

	h.serveEmbedded(w, r)
}

// serveEmbedded serves static files from the embedded [fs.FS] with SPA fallback.
func (h *SPAHandler) serveEmbedded(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		path = "index.html"
	}

	file, err := h.staticFS.Open(path)
	if err != nil {
		h.serveIndexHTML(w)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil || stat.IsDir() {
		h.serveIndexHTML(w)
		return
	}

	contentType := h.detectContentType(path)
	w.Header().Set("Content-Type", contentType)

	if h.isStaticAsset(path) {
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	} else {
		w.Header().Set("Cache-Control", "no-cache")
	}

	// Use ReadSeeker for proper range requests and caching
	if rs, ok := file.(io.ReadSeeker); ok {
		http.ServeContent(w, r, path, stat.ModTime(), rs)
	} else {
		_, _ = io.Copy(w, file)
	}
}

// serveIndexHTML serves the SPA index.html for client-side routing.
func (h *SPAHandler) serveIndexHTML(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	_, _ = w.Write(h.indexHTML) // #nosec G705 -- embedded static content
}

// detectContentType returns the MIME type for a file path.
func (h *SPAHandler) detectContentType(path string) string {
	ext := filepath.Ext(path)
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		return "application/octet-stream"
	}
	return contentType
}

// isStaticAsset determines if a file should be cached immutably.
func (h *SPAHandler) isStaticAsset(path string) bool {
	ext := filepath.Ext(path)
	staticExts := map[string]bool{
		".js":    true,
		".mjs":   true,
		".css":   true,
		".woff":  true,
		".woff2": true,
		".ttf":   true,
		".eot":   true,
		".otf":   true,
		".png":   true,
		".jpg":   true,
		".jpeg":  true,
		".gif":   true,
		".svg":   true,
		".ico":   true,
		".webp":  true,
		".avif":  true,
		".map":   true,
		".webm":  true,
		".mp4":   true,
		".mp3":   true,
		".wav":   true,
		".pdf":   true,
		".wasm":  true,
	}
	return staticExts[ext]
}
