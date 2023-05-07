package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iamelDuderino/my-website/pkg/games"
	"github.com/iamelDuderino/my-website/pkg/homepage"
)

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/v0/home", homepage.Display)
	http.HandleFunc("/api/v0/bumpball", games.BumpBall)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
