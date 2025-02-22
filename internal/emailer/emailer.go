package emailer

import (
	"fmt"
	"net/smtp"

	"github.com/iamelDuderino/my-website/internal/secretmanager"
)

func Send(fromName, fromEmail, msg string) error {
	var (
		err     error
		host    = secretmanager.Getenv("SMTP_HOST")
		port    = secretmanager.Getenv("SMTP_PORT")
		from    = secretmanager.Getenv("SMTP_FROM")
		to      = []string{secretmanager.Getenv("SMTP_TO")}
		pw      = secretmanager.Getenv("SMTP_APP_PW")
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
