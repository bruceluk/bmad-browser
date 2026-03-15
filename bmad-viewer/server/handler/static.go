package handler

import (
	"io/fs"
	"net/http"
	"strings"
)

// NewStaticHandler creates an HTTP handler that serves embedded frontend files
// with SPA fallback (returns index.html for routes not matching static files).
func NewStaticHandler(embeddedFS fs.FS) http.HandlerFunc {
	sub, err := fs.Sub(embeddedFS, "dist")
	if err != nil {
		panic("failed to create sub filesystem: " + err.Error())
	}

	fileServer := http.FileServer(http.FS(sub))

	return func(w http.ResponseWriter, r *http.Request) {
		// Skip API routes
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}

		// Try to serve static file
		path := strings.TrimPrefix(r.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}

		_, err := fs.Stat(sub, path)
		if err != nil {
			// File not found: SPA fallback to index.html
			r.URL.Path = "/"
		}

		fileServer.ServeHTTP(w, r)
	}
}
