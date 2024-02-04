package app

import "net/http"

func (x *app) buildRoutes() {

	// UI
	x.Server.Mux.HandleFunc("/", x.UI.HomeHandler)
	x.Server.Mux.HandleFunc("/about", x.UI.AboutHandler)
	x.Server.Mux.HandleFunc("/skills", x.UI.SkillsHandler)
	x.Server.Mux.HandleFunc("/contact", x.UI.ContactHandler)

	// File Server
	fs := http.FileServer(http.Dir("./public/"))
	x.Server.Mux.Handle("/public/", http.StripPrefix("/public/", fs))
}
