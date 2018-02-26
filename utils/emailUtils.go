package utils

import (
	"net/smtp"
	"strings"
)

const (
	EmailTypeGmail = iota
)

type Email struct {
	Address  string
	Password string
	Subject  string
	Body     string
	Type     int
}

func SendEmail(mail Email, to ...string) error {
	addr := ""
	var auth smtp.Auth
	switch mail.Type {
	case EmailTypeGmail:
		addr = "smtp.gmail.com:587"
		auth = smtp.PlainAuth("", mail.Address, mail.Password, "smtp.gmail.com")
	}
	body := "From: " + mail.Address + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Subject: " + mail.Subject + "\n\n" +
		mail.Body
	return smtp.SendMail(addr, auth, mail.Address, to, []byte(body))
}
