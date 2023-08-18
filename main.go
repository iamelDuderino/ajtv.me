package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/iamelDuderino/my-website/ui/views"
)

var (
	indexView   *views.View
	aboutView   *views.View
	skillsView  *views.View
	gamesView   *views.View
	contactView *views.View
	css         template.CSS
)

func main() {
	setCSS()
	indexView = views.NewView("layout", "./ui/views/index.gohtml")
	aboutView = views.NewView("layout", "./ui/views/about.gohtml")
	skillsView = views.NewView("layout", "./ui/views/skills.gohtml")
	gamesView = views.NewView("layout", "./ui/views/games.gohtml")
	contactView = views.NewView("layout", "./ui/views/contact.gohtml")
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/about", handleAbout)
	mux.HandleFunc("/skills", handleSkills)
	mux.HandleFunc("/games", handleGames)
	mux.HandleFunc("/contact", handleContact)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

type page struct {
	H1   string
	H2   string
	H3   string
	P    string
	OL   []string
	UL   []string
	CSS  template.CSS
	JS   template.JS
	Data interface{}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	err := indexView.Render(w, &page{
		H1:  "welcome to andrewjtomko.me!",
		CSS: css,
	})
	if err != nil {
		log.Println(err)
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	bio := getBio()
	aboutView.Render(w, &page{
		CSS:  css,
		Data: *bio,
	})
}

func handleSkills(w http.ResponseWriter, r *http.Request) {
	skillsView.Render(w, &page{
		H1:  "Skills",
		CSS: css,
	})
}

func handleGames(w http.ResponseWriter, r *http.Request) {
	gamesView.Render(w, &page{
		H1:  "Games",
		CSS: css,
	})
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	contactView.Render(w, &page{
		H1:  "Contact",
		CSS: css,
	})
}

func setCSS() {
	b, err := os.ReadFile("./ui/styles.css")
	if err != nil {
		panic(err)
	}
	css = template.CSS(string(b))
}

type resume struct {
	Summary   string
	Job1      *job
	Job2      *job
	Job3      *job
	Education []*edu
}

type job struct {
	CompanyName string
	Title       string
	Experience  string
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

func getBio() *bio {
	bio := &bio{
		FirstName:     "Andrew",
		LastName:      "Tomko",
		PreferredName: "AJ",
		Suffix:        "V",
	}
	bio.Resume.Summary = `Quick learner with a strong work ethic experienced in fast-paced onprem and cloud system administration from Active Directory and Cisco Unified Communications to G Suite, Azure AD, Zoom, WebEx and other SaaS applications with a mindset for security and a passion for building automation and process improvement including but not limited to synchronizing platforms through Powershell API scripting and SSO/SAML integrations.`
	bio.Resume.Job1 = &job{
		CompanyName: "Grafana Labs",
		Title:       "Enterprise Application Developer",
		Experience:  "Creating GoLang web apps and other cool stuff!",
		Years:       "02-14-22 thru Current",
	}
	bio.Resume.Job2 = &job{
		CompanyName: "Turbonomic, an IBM Company",
		Title:       "Manager, Global Help Desk",
		Experience:  "Managed a global help desk team",
		Years:       "01-02-2006 thru 02-14-22",
	}
	bio.Resume.Job3 = &job{
		CompanyName: "SevOne",
		Title:       "Sr IT Support Engineer",
		Experience:  "Escalations, Integrations, Mergers & Acquisitions, etc.",
		Years:       "03-04-1991 thru 01-02-2006",
	}
	bio.Resume.Education = append(bio.Resume.Education, &edu{
		School:       "Self Taught",
		Years:        "1991-current",
		DegreeOrCert: "Ceritification in Confidence",
	})
	return bio
}
