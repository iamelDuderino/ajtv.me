package app

import (
	"flag"

	utils "github.com/iamelDuderino/my-website/internal"
	"github.com/iamelDuderino/my-website/src/ui"
	"github.com/joho/godotenv"
)

const (
	requestType = "[APP]"
)

var Run app

type app struct {
	UI     *ui.UI
	Server httpServer
}

func init() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	utils.Logger.Log(requestType, utils.StatusProcessing, ".env loaded")

	// Runtime Flags
	dev := flag.Bool("dev", false, "Load local .env")
	flag.Parse()

	Run = app{
		UI: &ui.UI{},
	}

	Run.buildServer()
	Run.buildRoutes()
	Run.UI.BuildViews()
	Run.UI.BuildCookieStores(*dev)

}
