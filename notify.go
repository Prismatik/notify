package main

import (
	"github.com/prismatik/dotenv_safe"
	"github.com/prismatik/notify/email"
	"github.com/prismatik/notify/types"
	"os"
)

func init() {
	dotenv_safe.Load()
	switch os.Getenv("NOTIFY_EMAIL_PROVIDER") {
	case "gmail":
		dotenv_safe.LoadMany(dotenv_safe.Config{
			Envs:     []string{},
			Examples: []string{"example.gmail.env"},
		})
	}
}

func main() {
}
