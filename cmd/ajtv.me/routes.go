package main

import (
	"net/http"
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
	mux.Handle("GET /static/", http.StripPrefix("/static/", static))
	mux.Handle("GET /wasm/", http.StripPrefix("/wasm/", wasm))

	return mux
}
