package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/iamelDuderino/my-website/ui/views"
	"github.com/joho/godotenv"
)

var (
	homeView    *views.View
	aboutView   *views.View
	skillsView  *views.View
	gamesView   *views.View
	contactView *views.View
	css         template.CSS
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
		P:   "GoLang, Python, Powershell, HTML, CSS, JavaScript.. Okta, FreshService & BetterCloud Workflows.. Azure Web & Function App Deployments.. Building, Integrating & Maintaining APIs & Webhook Endpoints.. Slack Bots & Slash Commands.. and more!",
		CSS: css,
	})
}

// handleGames will be a grid layout with images of some simple sample projects
// that I started in JS in 2021 using SoloLearn, however, they will be refactored
// into an Ebiten application
func handleGames(w http.ResponseWriter, r *http.Request) {
	gamesView.Render(w, &page{
		H1:  "Games",
		P:   "Bump Ball | Pocket Pet Arena | Apex Legend Picker",
		CSS: css,
	})
}

// handleContact will present the Thank You page first if form has been submit
// otherwise it will present the contact form
func handleContact(w http.ResponseWriter, r *http.Request) {
	cname := r.FormValue("cname")
	cmsg := r.FormValue("cmsg")

	if cname != "" && cmsg != "" {
		contactView.Render(w, &page{
			H1:         "Thank You, " + cname + "!",
			P:          "I appreciate you reaching out and will respond as soon as possible!",
			CSS:        css,
			FormSubmit: true,
		})
		go sendMsg(cname, cmsg) // a go routine so that the page is not held up during signaling
		return
	}

	contactView.Render(w, &page{
		H1:         "Contact",
		P:          "Fill out the form below to send me an e-mail!",
		CSS:        css,
		FormSubmit: false,
	})
}

// setCSS saved the css file into the main reference for global use in templates
func setCSS() {
	b, err := os.ReadFile("./ui/styles.css")
	if err != nil {
		panic(err)
	}
	css = template.CSS(string(b))
}

// getBio is where a resume can be populated!
func getBio() *bio {
	bio := &bio{
		FirstName:     "Andrew",
		LastName:      "Tomko",
		PreferredName: "AJ",
		Suffix:        "V",
	}
	bio.Resume.Summary = `Quick, self-taught learner with a strong work ethic experienced in fast-paced on-prem/cloud, front-end/back-end system administration from Active Directory and Cisco Unified Communications to Okta, G Suite, Azure AD & Office 365 and many other SaaS applications with a mindset for security and a passion for building automation and process improvement tools including but not limited to synchronizing platforms with GoLang, Python and/or Powershell API scripting.`
	bio.Resume.Jobs = append(bio.Resume.Jobs, &job{
		CompanyName: "Grafana Labs",
		Title:       "Software Engineer I",
		Experience: []string{
			"Daily/Weekly/Quarterly Cron Jobs in Python for Auditing & Reporting utilizing API calls to query the various cloud services for MFA, Admin Role Additions/Removals, Okta-to-Slack Channel Syncs and more",
			"Designed, built, and maintained internal variant of Docebo Connect for Google Calendar & Zoom API/Webhook endpoint in GoLang which synchronizes employee LMS Sessions w/ Google Calendar, including synchronized Instructors-to-AltHosts with post-session recording url sharing via Slack",
			"Advanced Okta Workflow 'Flowgramming' utilizing Built-In & Custom API Connections with Helper Flows, Tables & Crons",
			"Maintained Slack Bot with /slash Command interactivity for end-user channel conversions with approval processes and quick and easy access to information from the various cloud services via Golang Web App backend",
			"Monitoring, Logging & Alerting implementations with Grafana Dashboards for visibility into enterprise application automations & integrations including but not limited to monitoring Workday Hires-to-date, Inbound New Hires over 90 days and cloud service license capacity",
			"Provisioned and maintained SAML/SSO/OIDC integrations for various cloud services with Okta Identity Engine and Azure Portal",
			"Coordinated the IT Operations efforts for BambooHR to Workday migration to account for any and all downstream impact and necessary maintenance window migrations from simple field updates to okta group rules and roles automations to custom internal integrations which utilized previous HRIS",
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

// sendMsg uses Environment Variables to send an Email using Gmail SMTP Servers
func sendMsg(name, msg string) error {
	var (
		err     error
		host    = os.Getenv("SMTP_HOST")
		port    = os.Getenv("SMTP_PORT")
		from    = os.Getenv("SMTP_FROM")
		to      = []string{os.Getenv("SMTP_TO")}
		pw      = os.Getenv("SMTP_APP_PW")
		subject = "You Have A New Message From " + name + "!"
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

	// Pages & UI
	setCSS()
	homeView = views.NewView("layout", "./ui/views/home.gohtml")
	aboutView = views.NewView("layout", "./ui/views/about.gohtml")
	skillsView = views.NewView("layout", "./ui/views/skills.gohtml")
	gamesView = views.NewView("layout", "./ui/views/games.gohtml")
	contactView = views.NewView("layout", "./ui/views/contact.gohtml")

	// Web Server & Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/about", handleAbout)
	mux.HandleFunc("/skills", handleSkills)
	mux.HandleFunc("/games", handleGames)
	mux.HandleFunc("/contact", handleContact)

	// Sample API (todo)

	// Sample Login Page (todo)

	// Static File Server
	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Listen & Serve
	log.Fatal(http.ListenAndServe(":8080", mux))

}
