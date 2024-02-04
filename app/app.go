package app

import (
	"flag"

	"github.com/iamelDuderino/my-website/internal/utils"
	"github.com/iamelDuderino/my-website/src/api"
	"github.com/iamelDuderino/my-website/src/ui"
	"github.com/joho/godotenv"
)

const (
	requestType = "[APP]"
)

var Run app

type app struct {
	UI     ui.UI
	API    api.API
	Server httpServer
}

func init() {

	utils.Logger.Log(requestType, utils.StatusBuilding, "App initialization started")

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	utils.Logger.Log(requestType, utils.StatusBuilding, ".env loaded")

	// Runtime Flags
	dev := flag.Bool("dev", false, "Load local .env")
	flag.Parse()

	// App Config
	Run.buildServer()
	utils.Logger.Log(requestType, utils.StatusBuilding, "app.buildServer complete")

	Run.buildRoutes()
	utils.Logger.Log(requestType, utils.StatusBuilding, "app.buildRoutes complete")

	// UI Config
	Run.UI.BuildViews()
	utils.Logger.Log(requestType, utils.StatusBuilding, "app.UI.BuildViews complete")

	Run.UI.BuildCookieStores(*dev)
	utils.Logger.Log(requestType, utils.StatusBuilding, "app.UI.BuildCookieStores complete")

	// App Ready To Run
	utils.Logger.Log(requestType, utils.StatusComplete, "App initialization complete")

}
