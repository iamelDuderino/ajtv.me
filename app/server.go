package app

import "net/http"

type httpServer struct {
	ListenAddress string
	Mux           *http.ServeMux
}

func (x *app) buildServer() {
	x.Server.ListenAddress = ":8080"
	x.Server.Mux = http.NewServeMux()
}
