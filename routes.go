package main

import (
	"net/http"
)

func (x *application) buildRoutes() {

	// UI
	x.server.mux.HandleFunc("/", x.ui.sessionManager(x.ui.home))
	x.server.mux.HandleFunc("/about", x.ui.sessionManager(x.ui.about))
	x.server.mux.HandleFunc("/skills", x.ui.sessionManager(x.ui.skills))
	x.server.mux.HandleFunc("/contact", x.ui.sessionManager(x.ui.contact))

	// API
	x.server.mux.HandleFunc("/api", x.api.getBasicResponse)

	// File Server
	fs := http.FileServer(http.Dir("./ui/static"))
	x.server.mux.Handle("/ui/static/", http.StripPrefix("/ui/static/", fs))

}
