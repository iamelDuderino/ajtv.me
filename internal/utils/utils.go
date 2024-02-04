package utils

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"
)

type Resume struct {
	Summary   string
	Jobs      []*Job
	Education []*Edu
}

type Job struct {
	CompanyName string
	Title       string
	Experience  []string
	StartDate   string
	EndDate     string
	Years       string
}

type Edu struct {
	School       string
	DegreeOrCert string
	Years        string
}

type Bio struct {
	FirstName     string
	LastName      string
	PreferredName string
	Suffix        string
	Resume        Resume
}

func SendEmail(fromName, fromEmail, msg string) error {
	var (
		err     error
		host    = os.Getenv("SMTP_HOST")
		port    = os.Getenv("SMTP_PORT")
		from    = os.Getenv("SMTP_FROM")
		to      = []string{os.Getenv("SMTP_TO")}
		pw      = os.Getenv("SMTP_APP_PW")
		subject = "You Have A New Message From " + fromName + " <" + fromEmail + ">!"
		b       = []byte(fmt.Sprintf("Subject: %s\n\n%s", subject, msg))
		auth    = smtp.PlainAuth(
			"",
			from,
			pw,
			host,
		)
	)
	err = smtp.SendMail(host+":"+port, auth, from, to, b)
	if err != nil {
		return err
	}
	return nil
}

// getBio reads a local json formatted resume
func GetBio() *Bio {
	b, err := os.ReadFile("./resume.json")
	if err != nil {
		return nil
	}
	bio := &Bio{}
	err = json.Unmarshal(b, &bio)
	if err != nil {
		return nil
	}
	return bio
}
