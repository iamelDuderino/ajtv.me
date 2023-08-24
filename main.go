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
	indexView = views.NewView("layout", "./ui/views/home.gohtml")
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
		P:   "GoLang, Python, Powershell, HTML, CSS, JavaScript.. Okta, FreshService & BetterCloud Workflows.. Azure Web & Function App Deployments.. Building, Integrating & Maintaining APIs & Webhook Endpoints.. Slack Bots & Slash Commands.. and more!",
		CSS: css,
	})
}

func handleGames(w http.ResponseWriter, r *http.Request) {
	gamesView.Render(w, &page{
		H1:  "Games",
		P:   "Bump Ball | Pocket Pet Arena | Apex Legend Picker",
		CSS: css,
	})
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	contactView.Render(w, &page{
		H1:  "Contact",
		P:   "Fill out the fake form from the future below to contact me!",
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

func getBio() *bio {
	bio := &bio{
		FirstName:     "Andrew",
		LastName:      "Tomko",
		PreferredName: "AJ",
		Suffix:        "V",
	}
	bio.Resume.Summary = `Quick learner with a strong work ethic experienced in fast-paced onprem and cloud system administration from Active Directory and Cisco Unified Communications to G Suite, Azure AD, Zoom, WebEx and other SaaS applications with a mindset for security and a passion for building automation and process improvement including but not limited to synchronizing platforms through GoLang API scripting and SSO/SAML integrations.`
	bio.Resume.Jobs = append(bio.Resume.Jobs, &job{
		CompanyName: "Grafana Labs",
		Title:       "Software Engineer I",
		Experience: []string{
			"Daily/Weekly/Quarterly Cron Jobs in Python for Auditing & Reporting utilizing API calls to query the various cloud services for MFA, Admin Role Additions/Removals, Okta-to-Slack Channel Syncs and more",
			"Designed, built, and maintained homebrew GoLang Docebo Connect for Google Calendar & Zoom which synchronizes employee LMS Sessions w/ Google Calendar, including synchronized Instructors-to-AltHosts with post-session recording url sharing via Slack",
			"Advanced Okta Workflow 'Flowgramming' utilizing Built-In & Custom API Connections with Helper Flows, Tables & Crons",
			"Monitoring, Logging & Alerting implementations with Grafana Dashboards for visibility into enterprise application automations & integrations including but not limited to monitoring Hires-to-date, Inbound Hires over 90 days and cloud service license capacity",
		},
		Years: "2022 - Current",
	})
	bio.Resume.Jobs = append(bio.Resume.Jobs, &job{
		CompanyName: "Turbonomic (acquired by IBM)",
		Title:       "Manager, Global Help Desk Services",
		Years:       "2021 - 2022",
		Experience: []string{
			"Responsible for Help Desk M&A activity post acquisition (IBM) in preparation for transfer of employment (TOE)",
			"Supported the transfer of multi-departmental services for critical business continuity in preparation for transfer of business (TOB)",
			"Coordinated the replacement of 520 employee assets for US, CAN and APAC",
			"Worked directly with Executives, Legal, and Human Resources sponsors to offer Keep Your Old Device (KYOD) programs for all Turbonomic assets to supplement Employee transitional concern and satisfaction",
		},
	})
	bio.Resume.Jobs = append(bio.Resume.Jobs, &job{
		CompanyName: "SevOne (acquired by Turbonomic)",
		Title:       "Team Lead, Help Desk",
		Years:       "2012 - 2021",
		Experience: []string{
			"Front and Back-end systems administration for Azure and Active Directories and other help desk managed industry standard SaaS applications",
			"Built and maintained multifunctional cross-company powershell tool utilizing API with CSV/JSON for routine help desk functions to create a single point of contact which decreased overall lead times of standard day-to-day requests and removed general human error",
			"Led the integration projects during acquisitions for migration of corporate devices, distribution lists, accounts and their relative SSO and other integrations",
			"Manage vendors for IT and Help Desk from initial discovery through the entire procurement process, including but not limited to, acquisition of hardware, software, cloud services, renewals and new vendor onboarding",
			"Successfully rolled out new Antivirus (Sophos) and MDM (workspace one)",
			"Supported Information Security in a successful ISO27001 implementation and audit",
			"Create and maintain all IT policies, procedures, standards and methodologies for the Help Desk team company wide",
		},
	})
	bio.Resume.Education = append(bio.Resume.Education, &edu{
		School:       "Self Taught",
		Years:        "1991-current",
		DegreeOrCert: "Ceritification in Confidence",
	})
	return bio
}
