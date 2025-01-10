package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var app application

type application struct {
	ui     userInterface
	api    applicationInterface
	server httpServer
}

func main() {
	flag.Parse()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	app.newServeMux()
	app.setRoutes()
	app.ui.buildViews()
	app.ui.buildCookieStores()
	log.Fatal(http.ListenAndServe(app.server.listenAddress, app.server.mux))
}
