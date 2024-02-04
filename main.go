package main

import (
	"log"
	"net/http"

	"github.com/iamelDuderino/my-website/app"
)

func main() {

	// download resume
	// uri := "https://" + os.Getenv("GITHUB_TOKEN") + "@raw.githubusercontent.com/" + os.Getenv("PATH_TO_PRIVATE_REPO")
	// req, err := http.NewRequest(http.MethodGet, uri, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// out, err := os.Create("resume.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = io.Copy(out, resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// resp.Body.Close()

	log.Fatal(http.ListenAndServe(app.Run.Server.ListenAddress, app.Run.Server.Mux))

}
