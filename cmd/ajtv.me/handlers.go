package main

import (
	"fmt"
	"net/http"

	"github.com/iamelDuderino/my-website/internal/emailer"
)

func (x *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	p := x.ui.newPage(r)
	err := x.ui.views["home"].render(w, p)
	if err != nil {
		x.logger.Error.Println(err)
	}
}

func (x *application) games(w http.ResponseWriter, r *http.Request) {
	p := x.ui.newPage(r)
	x.ui.views["games"].render(w, p)
}

func (x *application) blockbasher(w http.ResponseWriter, r *http.Request) {
	p := x.ui.newPage(r)
	err := x.ui.views["blockbasher"].render(w, p)
	if err != nil {
		x.logger.Error.Println(err)
	}
}

func (x *application) contact(w http.ResponseWriter, r *http.Request) {
	var (
		form = x.ui.newContactForm(r.FormValue("name"), r.FormValue("email"), r.FormValue("msg"))
		p    = x.ui.newPage(r)
	)
	p.Data = form
	x.ui.views["contact"].render(w, p)
}

func (x *application) contactPOST(w http.ResponseWriter, r *http.Request) {
	var (
		form = x.ui.newContactForm(r.FormValue("name"), r.FormValue("email"), r.FormValue("msg"))
		p    = x.ui.newPage(r)
	)
	if !form.Valid() {
		p.FlashMessage = "Please fill out all fields."
		p.Data = form
		x.ui.views["contact"].render(w, p)
		return
	}
	c, err := x.ui.globalSession.Get(r, globalSessionCookieName)
	if err != nil {
		x.logger.Error.Println("Error getting session:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	c.AddFlash("Thank you for reaching out!")
	c.Save(r, w)
	go emailer.Send(form.Name, form.Email, form.Message)
	http.Redirect(w, r, "/contact", http.StatusFound)
}

func (x *application) getBasicResponse(w http.ResponseWriter, r *http.Request) {
	resp := x.api.newResponse(true, "Success!")
	fmt.Fprint(w, resp.JSON())
}
