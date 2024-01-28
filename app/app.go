package app

var Run app

type app struct {
	Server httpServer
}

func init() {
	Run.buildServer()
	Run.buildRoutes()
}
