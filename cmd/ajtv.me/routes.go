package main

import (
	"io/fs"
	"net/http"

	"github.com/iamelDuderino/my-website/ui"
)

func (x *application) routes() http.Handler {

	mux := http.NewServeMux()

	// UI
	mux.HandleFunc("GET /", x.ui.sessionManager(x.home))
	mux.HandleFunc("GET /contact", x.ui.sessionManager(x.contact))
	mux.HandleFunc("GET /games", x.ui.sessionManager(x.games))
	mux.HandleFunc("GET /games/blockbasher", x.ui.sessionManager(x.blockbasher))
	// mux.HandleFunc("GET /about", x.ui.sessionManager(x.about))

	mux.HandleFunc("POST /contact", x.ui.sessionManager(x.contactPOST))

	// API
	mux.HandleFunc("GET /api", x.getBasicResponse)

	// File Server
	efss, err := fs.Sub(ui.EFS, "static")
	if err != nil {
		panic(err)
	}
	sefs := http.FileServer(http.FS(efss))
	mux.Handle("GET /static/", http.StripPrefix("/static/", sefs))

	efss, err = fs.Sub(ui.EFS, "games")
	if err != nil {
		panic(err)
	}
	gefs := http.FileServer(http.FS(efss))
	mux.Handle("GET /wasm/", http.StripPrefix("/wasm/", gefs))

	return mux
}
