package ui

import (
	"html/template"

	"github.com/gorilla/sessions"
)

type UI struct {
	homeView      *view
	aboutView     *view
	skillsView    *view
	contactView   *view
	globalSession *sessions.CookieStore
}

type page struct {
	Authenticated bool
	Data          interface{}
	CSS           template.CSS
	JS            template.JS
}

type view struct {
	Template *template.Template
	Layout   string
}
