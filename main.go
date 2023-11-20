package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/iamelDuderino/my-website/ui/views"
	"github.com/joho/godotenv"
)

var (
	homeView           *views.View
	aboutView          *views.View
	skillsView         *views.View
	gamesView          *views.View
	contactView        *views.View
	apexLegendsView    *views.View
	blockbasherView    *views.View
	pokemonstadiumView *views.View
	css                template.CSS
)

type page struct {
	H1         string
	H2         string
	H3         string
	P          string
	OL         []string
	UL         []string
	CSS        template.CSS
	JS         template.JS
	Data       interface{}
	FormSubmit bool
}

type resume struct {
	Summary   string
	Jobs      []*job
	Education []*edu
}

type job struct {
	CompanyName string
	Title       string
	Experience  []string
	StartDate   string
	EndDate     string
	Years       string
}

type edu struct {
	School       string
	DegreeOrCert string
	Years        string
}

type bio struct {
	FirstName     string
	LastName      string
	PreferredName string
	Suffix        string
	Resume        resume
}

// setCSS saves the css file into the main reference for global use in templates
func setCSS() {
	b, err := os.ReadFile("./ui/styles.css")
	if err != nil {
		panic(err)
	}
	css = template.CSS(string(b))
}

// getBio reads a local json formatted resume
func getBio() *bio {
	b, err := os.ReadFile("./resume.json")
	if err != nil {
		return nil
	}
	bio := &bio{}
	err = json.Unmarshal(b, &bio)
	if err != nil {
		return nil
	}
	return bio
}

// sendMsg uses Environment Variables to send an Email using Gmail SMTP Servers
func sendMsg(name, email, msg string) error {
	var (
		err     error
		host    = os.Getenv("SMTP_HOST")
		port    = os.Getenv("SMTP_PORT")
		from    = os.Getenv("SMTP_FROM")
		to      = []string{os.Getenv("SMTP_TO")}
		pw      = os.Getenv("SMTP_APP_PW")
		subject = "You Have A New Message From " + name + " <" + email + ">!"
		b       = []byte(fmt.Sprintf("Subject: %s\n\n%s", subject, msg))
		auth    = smtp.PlainAuth(
			"",
			from,
			pw,
			host,
		)
	)
	err = smtp.SendMail(host+":"+port, auth, from, to, b)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	// Runtime Flags
	local := flag.Bool("local", false, "Load local .env")
	flag.Parse()

	// if running locally load .env file
	if *local {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
		fmt.Println(".env loaded")
	}

	// download resume
	uri := "https://" + os.Getenv("GITHUB_TOKEN") + "@raw.githubusercontent.com/" + os.Getenv("PATH_TO_PRIVATE_REPO")
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.Create("resume.json")
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	// Pages & UI
	setCSS()
	homeView = views.NewView("layout", "./ui/views/home.gohtml")
	aboutView = views.NewView("layout", "./ui/views/about.gohtml")
	skillsView = views.NewView("layout", "./ui/views/skills.gohtml")
	gamesView = views.NewView("layout", "./ui/views/games.gohtml")
	contactView = views.NewView("layout", "./ui/views/contact.gohtml")
	apexLegendsView = views.NewView("layout", "./ui/views/apexlegends.gohtml")
	blockbasherView = views.NewView("layout", "./ui/views/blockbasher.gohtml")
	pokemonstadiumView = views.NewView("layout", "./ui/views/pokemonstadium.gohtml")

	// Web Server & Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/about", handleAbout)
	mux.HandleFunc("/skills", handleSkills)
	mux.HandleFunc("/games", handleGames)
	mux.HandleFunc("/contact", handleContact)

	// games
	mux.HandleFunc("/games/blockbasher", runBlockBasher)
	mux.HandleFunc("/games/pokemonstadium", runPokemonStadium)

	// other projects
	mux.HandleFunc("/projects/apexlegends", runApexLegends)

	// Sample API (todo)
	mux.HandleFunc("/api/apexlegends/stats", GETApexLegendsStats)
	mux.HandleFunc("/api/apexlegends/maprotation", GETApexLegendsMapRotation)

	// Sample Login Page (todo)

	// Static File Server
	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// WASM File Server
	blockbasherWASM := http.FileServer(http.Dir("./games/blockbasher/wasm"))
	mux.Handle("/wasm/blockbasher/run.html", http.StripPrefix("/wasm/blockbasher/", blockbasherWASM))

	// Listen & Serve
	log.Fatal(http.ListenAndServe(":8080", mux))

}
