package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iamelDuderino/my-website/src/games"
	"github.com/iamelDuderino/my-website/src/homepage"
)

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/", homepage.Display)
	http.HandleFunc("/bumpball", games.BumpBall)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
