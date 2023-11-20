package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/iamelDuderino/my-website/games/apexlegends"
)

// handleHome is the home page!
func handleHome(w http.ResponseWriter, r *http.Request) {
	err := homeView.Render(w, &page{
		H1:  "Hello!",
		CSS: css,
	})
	if err != nil {
		log.Println(err)
	}
}

// handleAbout is a templated Resume layout that expands bullets as needed
func handleAbout(w http.ResponseWriter, r *http.Request) {
	bio := getBio()
	aboutView.Render(w, &page{
		CSS:  css,
		Data: *bio,
	})
}

// handleSkills is a simple skill page that should be prettied up
// with some fancier buttons/tags or something
func handleSkills(w http.ResponseWriter, r *http.Request) {
	skillsView.Render(w, &page{
		H1:  "Skills",
		P:   "Paid Problem Solver! GoLang, Python, Powershell, HTML, CSS, JavaScript.. Okta, FreshService & BetterCloud Workflows.. Azure Web & Function App Deployments.. Building, Integrating & Maintaining APIs & Webhook Endpoints.. Slack Bots & Slash Commands.. and more!",
		CSS: css,
	})
}

// handleGames will be a grid layout with images of some simple sample projects
// that I started in JS in 2021 using SoloLearn, however, they will be refactored
// into an Ebiten application
func handleGames(w http.ResponseWriter, r *http.Request) {
	type game struct {
		Name string
		Link string
	}
	type games struct {
		Games []*game
	}
	data := &games{}
	data.Games = append(data.Games, &game{
		Name: "Block Basher",
		Link: "/games/blockbasher",
	})
	data.Games = append(data.Games, &game{
		Name: "Pokemon Stadium",
		Link: "/games/pokemonstadium",
	})
	data.Games = append(data.Games, &game{
		Name: "Apex Legends",
		Link: "/projects/apexlegends",
	})

	gamesView.Render(w, &page{
		H1:   "Games & Projects",
		Data: data,
		CSS:  css,
	})
}

// handleContact will present the Thank You page first if form has been submit
// otherwise it will present the contact form
func handleContact(w http.ResponseWriter, r *http.Request) {
	cname := r.FormValue("cname")
	cmsg := r.FormValue("cmsg")
	cemail := r.FormValue("cemail")

	if cname != "" && cmsg != "" {
		contactView.Render(w, &page{
			H1:         "Thank You, " + cname + "!",
			P:          "I appreciate you reaching out!",
			CSS:        css,
			FormSubmit: true,
		})
		go sendMsg(cname, cemail, cmsg) // a go routine so that the page is not held up during signaling
		return
	}

	contactView.Render(w, &page{
		H1:         "Contact",
		P:          "Fill out the form below to send me an e-mail!",
		CSS:        css,
		FormSubmit: false,
	})
}

func runApexLegends(w http.ResponseWriter, r *http.Request) {
	player := r.FormValue("player")
	platform := r.FormValue("platform")
	mapRotation, _ := apexlegends.GetMapRotation()
	randomLegend := apexlegends.ApexLegendsConfig.RandomLegend()
	page := &page{
		CSS: css,
	}
	type extraData struct {
		MapRotation       *apexlegends.MapRotation
		RandomLegend      *apexlegends.Legend
		PlayerStats       *apexlegends.PlayerStats
		AdditionalContent []string
		Error             string
	}
	data := &extraData{
		MapRotation:  mapRotation,
		RandomLegend: randomLegend,
	}

	page.H1 = "Apex Legends"
	page.P = "This page will connect to https://apexlegendsapi.com to query player stats per platform, provide the current map rotation, or suggest your next legend to play at random!"
	page.Data = data

	if player != "" && platform != "" {
		stats, err := apexlegends.GetPlayerStats(player, platform)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		data.PlayerStats = stats
		if stats.Error != "" {
			data.Error = stats.Error
		}
		page.FormSubmit = true
	}

	apexLegendsView.Render(w, page)

}

func runPokemonStadium(w http.ResponseWriter, r *http.Request) {
	pokemonstadiumView.Render(w, &page{
		H1:  "Pokemon Stadium",
		P:   "This page is still under construction.. Please come back later!",
		CSS: css,
	})
}

func runBlockBasher(w http.ResponseWriter, r *http.Request) {
	blockbasherView.Render(w, &page{
		CSS: css,
	})
}

func GETApexLegendsStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, `Unsupported Method`)
		return
	}
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	var (
		response         []byte
		player, platform string
	)
	if player = r.Form.Get("player"); player == "" {
		fmt.Fprint(w, `player not provided`)
		return
	}
	if platform = r.Form.Get("platform"); platform == "" {
		fmt.Fprint(w, `platform not provided`)
		return
	}
	stats, err := apexlegends.GetPlayerStats(player, platform)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	response, err = json.Marshal(stats)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	w.Write(response)
}

func GETApexLegendsMapRotation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, `Unsupported Method`)
		return
	}
	var (
		response []byte
	)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	mapr, err := apexlegends.GetMapRotation()
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	response, err = json.Marshal(mapr)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	w.Write(response)
}
