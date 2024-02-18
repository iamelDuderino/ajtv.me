package main

import (
	"fmt"
	"net/http"

	"github.com/iamelDuderino/my-website/internal/utils"
)

func (x *userInterface) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	p := x.newPage(r)
	x.homeView.render(w, p)
}

// sample simple admin page
// func (x *userInterface) admin(w http.ResponseWriter, r *http.Request) {
// 	p := x.newPage(r)
// 	if !p.Authenticated {
// 		http.Redirect(w, r, "/", http.StatusOK)
// 		return
// 	}
// 	x.homeView.render(w, p)
// }

func (x *userInterface) about(w http.ResponseWriter, r *http.Request) {
	bio := utils.ReadResume()
	p := x.newPage(r)
	p.Data = bio
	x.aboutView.render(w, p)
}

func (x *userInterface) skills(w http.ResponseWriter, r *http.Request) {
	p := x.newPage(r)
	x.skillsView.render(w, p)
}

func (x *userInterface) contact(w http.ResponseWriter, r *http.Request) {
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

func (x *applicationInterface) getBasicResponse(w http.ResponseWriter, r *http.Request) {
	resp := x.newResponse()
	resp.Message = "Thank you for testing my basic sample API!"
	fmt.Fprint(w, resp.JSON())
}
