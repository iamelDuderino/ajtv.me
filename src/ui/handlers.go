package ui

import (
	"net/http"

	"github.com/iamelDuderino/my-website/internal/utils"
)

func (x *UI) HomeHandler(w http.ResponseWriter, r *http.Request) {
	p := x.newPage(r)
	x.HomeView.Render(w, p)
}

// handleAbout is a templated Resume layout that expands bullets as needed
func (x *UI) AboutHandler(w http.ResponseWriter, r *http.Request) {
	bio := utils.GetBio()
	p := x.newPage(r)
	p.Data = bio
	x.AboutView.Render(w, p)
}

// handleSkills is a simple skill page that should be prettied up
// with some fancier buttons/tags or something
func (x *UI) SkillsHandler(w http.ResponseWriter, r *http.Request) {
	p := x.newPage(r)
	x.SkillsView.Render(w, p)
	// 	H1:  "Skills",
	// 	P:   "Paid Problem Solver! GoLang, Python, Powershell, HTML, CSS, JavaScript.. Okta, FreshService & BetterCloud Workflows.. Azure Web & Function App Deployments.. Building, Integrating & Maintaining APIs & Webhook Endpoints.. Slack Bots & Slash Commands.. and more!",
	// 	CSS: css,
	// })
}

// handleContact will present the Thank You page first if form has been submit
// otherwise it will present the contact form
func (x *UI) ContactHandler(w http.ResponseWriter, r *http.Request) {
	var (
		cname  = r.FormValue("cname")
		cemail = r.FormValue("cemail")
		cmsg   = r.FormValue("cmsg")
		p      = x.newPage(r)
	)
	if cname != "" && cemail != "" && cmsg != "" {
		p.Data = true
		go utils.SendEmail(cname, cemail, cmsg)
	}
	x.ContactView.Render(w, p)
}
