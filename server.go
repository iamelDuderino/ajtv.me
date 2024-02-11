package main

import "net/http"

type httpServer struct {
	listenAddress string
	mux           *http.ServeMux
}

func (x *application) newServeMux() {
	x.server.listenAddress = ":8080"
	x.server.mux = http.NewServeMux()
}
