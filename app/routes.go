package app

import (
	"net/http"
)

func (x *app) buildRoutes() {

	// UI
	x.Server.Mux.HandleFunc("/", x.UI.SessionManager(x.UI.HomeHandler))
	x.Server.Mux.HandleFunc("/about", x.UI.SessionManager(x.UI.AboutHandler))
	x.Server.Mux.HandleFunc("/skills", x.UI.SessionManager(x.UI.SkillsHandler))
	x.Server.Mux.HandleFunc("/contact", x.UI.SessionManager(x.UI.ContactHandler))

	// API
	x.Server.Mux.HandleFunc("/api", x.API.HomeHandler)

	// File Server
	fs := http.FileServer(http.Dir("./public/"))
	x.Server.Mux.Handle("/public/", http.StripPrefix("/public/", fs))

}
