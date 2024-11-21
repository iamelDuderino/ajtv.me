package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/sessions"
	"github.com/iamelDuderino/my-website/internal/utils"
)

const (
	viewsFolder             = "./ui/views"
	globalSessionCookieName = "ajtv.me-global-cookies"
)

var (
	css template.CSS
)

type page struct {
	Authenticated bool
	Data          interface{}
	FlashMessage  string
	CSS           template.CSS
	JS            template.JS
}

func (x *page) GetURL(url string) string {
	var (
		accepted = []string{
			"GITHUB",
			"LINKEDIN",
		}
		allowed = func(s string) bool {
			for _, i := range accepted {
				if i == s {
					return true
				}
			}
			return false
		}
	)
	if !allowed(url) {
		return ""
	}
	return os.Getenv(url + "_URL")
}

type view struct {
	Template *template.Template
	Layout   string
}

func (v *view) render(w http.ResponseWriter, data interface{}) error {
	var buf bytes.Buffer
	err := v.Template.ExecuteTemplate(&buf, v.Layout, data)
	if err != nil {
		return err
	}
	fmt.Fprint(w, buf.String())
	return nil
}

type userInterface struct {
	homeView      *view
	aboutView     *view
	skillsView    *view
	contactView   *view
	globalSession *sessions.CookieStore
}

func (x *userInterface) buildViews() {

	// Views
	x.homeView = x.newView("base", viewsFolder+"/home.gohtml")
	x.aboutView = x.newView("base", viewsFolder+"/about.gohtml")
	x.skillsView = x.newView("base", viewsFolder+"/skills.gohtml")
	x.contactView = x.newView("base", viewsFolder+"/contact.gohtml")

	// CSS
	b, err := os.ReadFile("./ui/styles.css")
	if err != nil {
		panic(err)
	}
	css = template.CSS(string(b))

}

func (x *userInterface) buildCookieStores(dev bool) {
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

func (x *userInterface) newView(layout string, files ...string) *view {
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

func (x *userInterface) newPage(r *http.Request) *page {
	var p = &page{
		CSS: css,
	}

	// placeholder use r to set p.Authenticated if logged in with Okta Dev console
	// for sample simple admin page

	return p
}

func (x *userInterface) getTemplateFiles() []string {
	files, err := filepath.Glob("./ui/templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

func (x *userInterface) sessionManager(fn func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
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

func (x *userInterface) newContactForm(name, email, msg string) *ContactForm {
	return &ContactForm{
		Name:    name,
		Email:   email,
		Message: msg,
		Errors:  make(map[string]string),
		Visible: true,
	}
}

type ContactForm struct {
	Name, Email, Message string
	Visible              bool
	Errors               map[string]string
}

func (x *ContactForm) isValid() bool {
	validName := x.validName()
	if !validName {
		x.Errors["Name"] = "Name Is Blank!"
	}
	validEmail := x.validEmail()
	if !validEmail {
		x.Errors["Email"] = "Email Is Blank!"
	}
	validMessage := x.validMessage()
	if !validMessage {
		x.Errors["Message"] = "Message Is Blank!"
	}
	return (len(x.Errors) == 0)
}

func (x *ContactForm) validName() bool {
	return (x.Name != "" && x.Name != " ")
}

func (x *ContactForm) validEmail() bool {
	return (x.Email != "" && x.Email != " ")
}

func (x *ContactForm) validMessage() bool {
	return (x.Message != "" && x.Message != " ")
}

func (x *ContactForm) clear() {
	x.Name = ""
	x.Email = ""
	x.Message = ""
}
