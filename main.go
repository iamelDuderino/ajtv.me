package main

import (
	"log"
	"net/http"

	"github.com/iamelDuderino/my-website/app"
)

func main() {
	log.Fatal(http.ListenAndServe(app.Run.Server.ListenAddress, app.Run.Server.Mux))
}
