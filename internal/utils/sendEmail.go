package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(fromName, fromEmail, msg string) error {
	var (
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
	err := smtp.SendMail(host+":"+port, auth, from, to, b)
	if err != nil {
		return err
	}
	return nil
}
