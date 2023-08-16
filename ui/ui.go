package ui

import (
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

type page struct {
	Title string
	Data  *data
	CSS   template.CSS
	JS    template.JS
}

type data struct {
	H1 string
	H2 string
	H3 string
	P  string
	OL []string
	UL []string
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	p := &page{
		Title: "home",
		Data: &data{
			H1: "Home Page",
		},
	}
	t, e := t.ParseFiles("./www/templates/home.gohtml")
	if e != nil {
		fmt.Fprint(w, e.Error())
		return
	}
	e = t.Execute(w, p)
	if e != nil {
		fmt.Fprint(w, e.Error())
		return
	}
}

// type resumeData struct {
// 	CompanyName     string
// 	JobTitle        string
// 	EmploymentBegan string
// 	EmploymentEnded string
// }

func Resume(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	p := &page{
		Title: "resume",
		Data: &data{
			H1: "Andrew J Tomko V",
			H2: "Sr. IT Operations Engineer",
			P:  "Quick learner with a strong work ethic",
		},
	}
	t, e := t.ParseFiles("./www/templates/resume.gohtml")
	if e != nil {
		fmt.Fprint(w, e.Error())
		return
	}
	e = t.Execute(w, p)
	if e != nil {
		fmt.Fprint(w, e.Error())
		return
	}
}
