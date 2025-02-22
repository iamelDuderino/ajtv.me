package main

import (
	"compress/gzip"
	"io/fs"
	"net/http"
	"strings"

	"github.com/iamelDuderino/my-website/ui"
)

func (x *application) routes() http.Handler {

	mux := http.NewServeMux()

	// UI
	mux.HandleFunc("GET /", x.ui.sessionManager(x.home))
	mux.HandleFunc("GET /contact", x.ui.sessionManager(x.contact))
	mux.HandleFunc("GET /games", x.ui.sessionManager(x.games))
	mux.HandleFunc("GET /games/blockbasher", x.ui.sessionManager(x.blockbasher))

	mux.HandleFunc("POST /contact", x.ui.sessionManager(x.contactPOST))

	// API
	mux.HandleFunc("GET /api", x.getBasicResponse)

	// File Servers
	efsDir, err := fs.Sub(ui.EFS, "static")
	if err != nil {
		panic(err)
	}
	staticEFS := http.FileServer(http.FS(efsDir))
	mux.Handle("GET /static/", http.StripPrefix("/static/", staticEFS))

	efsDir, err = fs.Sub(ui.EFS, "wasm")
	if err != nil {
		panic(err)
	}
	wasmEFS := http.FileServer(http.FS(efsDir))
	wasmHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Serve actual .wasm files with gzip encoding
		if strings.HasSuffix(r.URL.Path, ".wasm") {
			w.Header().Set("Content-Type", "application/wasm")
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)
			defer gz.Close()
			gw := &gzipResponseWriter{ResponseWriter: w, Writer: gz}
			wasmEFS.ServeHTTP(gw, r)

		} else {
			// Serve the wasm html iframe template
			wasmEFS.ServeHTTP(w, r)
		}
	})
	mux.Handle("GET /wasm/", http.StripPrefix("/wasm/", wasmHandler))

	return mux
}
