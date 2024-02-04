package ui

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

const (
	requestType = "[UI]"

	viewsFolder = "./src/ui/views"
)

var (
	css template.CSS
)

type UI struct {
	HomeView    *view
	AboutView   *view
	SkillsView  *view
	ContactView *view
}

func (x *UI) BuildViews() {

	// utils.Logger.Log(requestType, utils.StatusBuilding, "ui.BuildViews started")

	x.HomeView = x.newView("base", viewsFolder+"/home.gohtml")
	x.AboutView = x.newView("base", viewsFolder+"/about.gohtml")
	x.SkillsView = x.newView("base", viewsFolder+"/skills.gohtml")
	x.ContactView = x.newView("base", viewsFolder+"/contact.gohtml")

	// set css
	b, err := os.ReadFile("./src/ui/styles.css")
	if err != nil {
		panic(err)
	}
	css = template.CSS(string(b))

}

func (x *UI) BuildCookieStores(dev bool) {
	var params string
	switch dev {
	case true:
		params = "localhost"
	case false:
		params = "gcloud.ue.r.appspot.com"
	}
	fmt.Println(requestType + params)
	// utils.Logger.Log(requestType, utils.StatusComplete, "ui.BuildCookieStores complete for "+params)
}

func (x *UI) newView(layout string, files ...string) *view {
	files = append(x.getTemplateFiles(), files...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &view{
		Template: t,
		Layout:   layout,
	}
}

func (x *UI) newPage(r *http.Request) *page {
	return &page{
		CSS: css,
	}
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

func (v *view) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func (x *UI) getTemplateFiles() []string {
	files, err := filepath.Glob("./src/ui/templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}
