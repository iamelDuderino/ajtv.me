package ui

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/sessions"
	"github.com/iamelDuderino/my-website/internal/utils"
)

func (x *UI) BuildViews() {

	// Views
	x.homeView = x.newView("base", viewsFolder+"/home.gohtml")
	x.aboutView = x.newView("base", viewsFolder+"/about.gohtml")
	x.skillsView = x.newView("base", viewsFolder+"/skills.gohtml")
	x.contactView = x.newView("base", viewsFolder+"/contact.gohtml")

	// CSS
	b, err := os.ReadFile("./src/ui/styles.css")
	if err != nil {
		panic(err)
	}
	css = template.CSS(string(b))

}

func (x *UI) BuildCookieStores(dev bool) {
	// placeholder content for gorilla sessions implementation
	// var params string
	// switch dev {
	// case true:
	// 	params = "localhost"
	// case false:
	// 	params = "gcloud.ue.r.appspot.com"
	// }
	x.globalSession = sessions.NewCookieStore([]byte(os.Getenv("GLOBAL_SESSION_SECRET")))
}

func (x *UI) newView(layout string, files ...string) *view {
	files = append(files, x.getTemplateFiles()...)
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

func (x *UI) getTemplateFiles() []string {
	files, err := filepath.Glob("./src/ui/templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

func (x *UI) SessionManager(fn func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s, err := x.globalSession.Get(r, globalSessionCookieName)
		if err != nil {
			utils.Logger.LogErr(requestType, err)
		}
		if s.IsNew {
			err = s.Save(r, w)
			if err != nil {
				utils.Logger.LogErr(requestType, err)
			}
		}
		fn(w, r)
	}
}

func (v *view) render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
