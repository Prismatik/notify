package gmail

import (
	"github.com/prismatik/notify/types"
	"net/smtp"
	"os"
)

func Send(e types.Email) error {
	from := os.Getenv("NOTIFY_EMAIL_FROM")
	pass := os.Getenv("NOTIFY_EMAIL_SMTP_PASS")

	msg := "Return-Path: " + from + "\n" +
		"From: " + e.From + "\n" +
		"To: " + e.To + "\n" +
		"Subject: " + e.Subject + "\n\n" +
		e.Body

	return smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{e.To}, []byte(msg))
}
