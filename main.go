package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/iamelDuderino/my-website/internal/utils"
	"github.com/joho/godotenv"
)

const (
	requestType = "[APP]"
)

var app application

type application struct {
	ui     userInterface
	api    applicationInterface
	server httpServer
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	utils.Logger.Log(requestType, utils.StatusBuilding, ".env loaded")

	// Runtime Flags
	dev := flag.Bool("dev", false, "Load local .env")
	flag.Parse()

	// Server Config
	app.newServeMux()
	utils.Logger.Log(requestType, utils.StatusBuilding, "app.newServeMux complete")

	app.buildRoutes()
	utils.Logger.Log(requestType, utils.StatusBuilding, "app.buildRoutes complete")

	// UI Config
	app.ui.buildViews()
	utils.Logger.Log(requestType, utils.StatusBuilding, "app.ui.buildViews complete")

	app.ui.buildCookieStores(*dev)
	utils.Logger.Log(requestType, utils.StatusBuilding, "app.ui.buildCookieStores complete")

	// App Ready To Run
	utils.Logger.Log(requestType, utils.StatusComplete, "app is ready!")
	log.Fatal(http.ListenAndServe(app.server.listenAddress, app.server.mux))
}
