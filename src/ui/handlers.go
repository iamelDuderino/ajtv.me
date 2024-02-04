package ui

import (
	"net/http"

	"github.com/iamelDuderino/my-website/internal/utils"
)

func (x *UI) HomeHandler(w http.ResponseWriter, r *http.Request) {
	p := x.newPage(r)
	x.homeView.render(w, p)
}

func (x *UI) AboutHandler(w http.ResponseWriter, r *http.Request) {
	bio := utils.ReadResume()
	p := x.newPage(r)
	p.Data = bio
	x.aboutView.render(w, p)
}

func (x *UI) SkillsHandler(w http.ResponseWriter, r *http.Request) {
	p := x.newPage(r)
	x.skillsView.render(w, p)
}

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
	x.contactView.render(w, p)
}
