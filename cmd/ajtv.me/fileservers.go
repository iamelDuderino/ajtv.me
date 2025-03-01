package main

import (
	"compress/gzip"
	"io/fs"
	"net/http"
	"strings"
	"time"

	"github.com/iamelDuderino/my-website/ui"
)

var static = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	dir, err := fs.Sub(ui.EFS, "static")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	w.Header().Set("Expires", time.Now().AddDate(1, 0, 0).Format(http.TimeFormat))
	efs := http.FileServer(http.FS(dir))
	efs.ServeHTTP(w, r)
})

var wasm = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	dir, err := fs.Sub(ui.EFS, "wasm")
	if err != nil {
		panic(err)
	}

	efs := http.FileServer(http.FS(dir))

	// Serve actual .wasm files with gzip encoding
	if strings.HasSuffix(r.URL.Path, ".wasm") {
		w.Header().Set("Content-Type", "application/wasm")
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gw := &gzipResponseWriter{ResponseWriter: w, Writer: gz}
		efs.ServeHTTP(gw, r)

	} else {
		// Serve the wasm html iframe template
		efs.ServeHTTP(w, r)
	}
})
