package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/iamelDuderino/my-website/internal/logger"
	"github.com/iamelDuderino/my-website/internal/secretmanager"
	"github.com/iamelDuderino/my-website/ui"
)

const (
	viewsFolder             = "views/"
	globalSessionCookieName = "ajtv.me-global-cookies"
)

var (
	css template.CSS
)

type page struct {
	Authenticated bool
	Data          any
	FlashMessage  string
	CSS           template.CSS
	JS            template.JS
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
	views         map[string]*view
	globalSession *sessions.CookieStore
	logger        *logger.Logger
}

func (x *userInterface) buildViews() {

	// Views
	x.views["home"] = x.newView("base", viewsFolder+"home.gohtml")
	x.views["about"] = x.newView("base", viewsFolder+"about.gohtml")
	x.views["skills"] = x.newView("base", viewsFolder+"skills.gohtml")
	x.views["contact"] = x.newView("base", viewsFolder+"contact.gohtml")
	x.views["games"] = x.newView("base", viewsFolder+"games.gohtml")
	x.views["blockbasher"] = x.newView("base", viewsFolder+"blockbasher.gohtml")

	// CSS
	css = template.CSS(ui.StyleSheet)

}

func (x *userInterface) buildCookieStores(dev bool) {
	x.globalSession = sessions.NewCookieStore([]byte(secretmanager.Getenv("GLOBAL_SESSION_SECRET")))
	// placeholder content for gorilla sessions implementation
	// switch dev {
	// case true:
	// 	x.globalSession.Options.Domain = "ajtv.me.local"
	// case false:
	// 	x.globalSession.Options.Domain = "ajtv.me"
	// }
}

func (x *userInterface) newView(layout string, files ...string) *view {
	f := x.getTemplateFiles()
	f = append(f, files...)
	t, err := template.New(layout).ParseFS(ui.EFS, f...)
	if err != nil {
		panic(err)
	}
	return &view{
		Template: t,
		Layout:   layout,
	}
}

func (x *userInterface) getTemplateFiles() []string {
	files, err := fs.Glob(ui.EFS, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

func (x *userInterface) newPage(r *http.Request) *page {
	var p = &page{
		CSS: css,
	}

	// placeholder use r to set p.Authenticated if logged in with Okta Dev console
	// for sample simple admin page

	return p
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

func (x *ContactForm) Valid() bool {
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
