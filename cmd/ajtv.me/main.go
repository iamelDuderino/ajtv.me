package main

import (
	"flag"
	"net/http"

	"github.com/iamelDuderino/my-website/internal/logger"
)

type application struct {
	ui     *userInterface
	api    *applicationInterface
	server *httpServer
	logger *logger.Logger
}

func newApplication(dev bool) *application {
	app := &application{
		ui: &userInterface{
			views:  make(map[string]*view),
			logger: logger.NewLogger("UI_INFO", "UI_ERROR"),
		},
		api:    new(applicationInterface),
		logger: logger.NewLogger("INFO", "ERROR"),
	}
	app.server = app.buildServer()
	app.ui.buildViews()
	app.ui.buildCookieStores(dev)
	return app
}

func main() {
	dev := flag.Bool("dev", false, "Load local development configuration for session cookies")
	flag.Parse()
	app := newApplication(*dev)
	app.logger.Info.Println("Starting Server on", app.server.listenAddress)
	app.logger.Error.Fatal(http.ListenAndServe(app.server.listenAddress, app.server.mux))
}
