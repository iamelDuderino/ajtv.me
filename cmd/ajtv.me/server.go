package main

import "net/http"

const (
	port = ":8080"
)

type httpServer struct {
	listenAddress string
	mux           http.Handler
}

func (x *application) buildServer() *httpServer {
	return &httpServer{
		listenAddress: port,
		mux:           x.routes(),
	}
}
