package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prismatik/dotenv_safe"
	"github.com/prismatik/notify/email"
	"github.com/prismatik/notify/sms"
	"github.com/prismatik/notify/types"
	"net/http"
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
	switch os.Getenv("NOTIFY_SMS_PROVIDER") {
	case "amazon":
		dotenv_safe.LoadMany(dotenv_safe.Config{
			Envs:     []string{},
			Examples: []string{"example.amazon_sms.env"},
		})
	}
}

func main() {
	r := gin.Default()

	r.POST("/sms", func(c *gin.Context) {
		var json types.SMS
		if c.BindJSON(&json) == nil {
			err := sms.Send(json)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			}
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	})

	r.POST("/email", func(c *gin.Context) {
		var json types.Email
		if c.BindJSON(&json) == nil {
			err := email.Send(json)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			}
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	})

	r.Run()
}
