package main

import (
	"net/http"

	"github.com/iamelDuderino/my-website/internal/logger"
)

type application struct {
	ui     *userInterface
	api    *applicationInterface
	server *httpServer
	logger *logger.Logger
}

func newApplication() *application {
	app := &application{
		ui: &userInterface{
			views:  make(map[string]*view),
			logger: logger.NewLogger("UI_INFO", "UI_ERROR"),
		},
		api: &applicationInterface{
			logger: logger.NewLogger("API_INFO", "API_ERROR"),
		},
		logger: logger.NewLogger("APP_INFO", "APP_ERROR"),
	}
	app.server = app.buildServer()
	app.ui.buildViews()
	app.ui.buildCookieStores()
	return app
}

func main() {
	app := newApplication()
	app.logger.Info.Println("Starting Server on", app.server.listenAddress)
	app.logger.Error.Fatal(http.ListenAndServe(app.server.listenAddress, app.server.mux))
}
