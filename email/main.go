package email

import (
	"errors"
	"github.com/prismatik/notify/email/gmail"
	"github.com/prismatik/notify/types"
	"os"
)

func Send(e types.Email) error {
	switch os.Getenv("NOTIFY_EMAIL_PROVIDER") {
	case "gmail":
		return gmail.Send(e)
	default:
		return errors.New("No valid email provider configured")
	}
}
