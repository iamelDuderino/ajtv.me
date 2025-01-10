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
	p := x.newPage(r)
	p.Data = utils.ReadResume()
	x.aboutView.render(w, p)
}

func (x *userInterface) contact(w http.ResponseWriter, r *http.Request) {
	var (
		form = x.newContactForm(r.FormValue("name"), r.FormValue("email"), r.FormValue("msg"))
		p    = x.newPage(r)
	)
	if r.Method == http.MethodPost {
		ok := form.valid()
		if ok {
			p.FlashMessage = "Thank you for reaching out!"
			form.clear()
			form.Visible = false
			go utils.SendEmail(form.Name, form.Email, form.Message)
		}
	}
	p.Data = form
	x.contactView.render(w, p)
}

func (x *applicationInterface) getBasicResponse(w http.ResponseWriter, r *http.Request) {
	resp := x.newResponse(true, "Success!")
	fmt.Fprint(w, resp.JSON())
}
