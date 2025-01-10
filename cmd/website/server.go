package main

import (
	"net/http"
)

type httpServer struct {
	listenAddress string
	mux           *http.ServeMux
}

func (x *application) newServeMux() {
	x.server.listenAddress = ":8080"
	x.server.mux = http.NewServeMux()
}

func (x *application) setRoutes() {

	// UI
	x.server.mux.HandleFunc("/", x.ui.sessionManager(x.ui.home))
	x.server.mux.HandleFunc("/about", x.ui.sessionManager(x.ui.about))
	x.server.mux.HandleFunc("/contact", x.ui.sessionManager(x.ui.contact))

	// API
	x.server.mux.HandleFunc("/api", x.api.get(x.api.getBasicResponse))

	// File Server
	fs := http.FileServer(http.Dir("./ui/static"))
	x.server.mux.Handle("/ui/static/", http.StripPrefix("/ui/static/", fs))

}
