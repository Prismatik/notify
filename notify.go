package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/prismatik/dotenv_safe"
	"github.com/prismatik/notify/email"
	"github.com/prismatik/notify/sms"
	"github.com/prismatik/notify/types"
	"log"
	"net/http"
	"os"
	"time"
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
	r := mux.NewRouter()

	r.HandleFunc("/sms", SmsHandler).
		Methods("POST")

	r.HandleFunc("/email", EmailHandler).
		Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func SmsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var t types.SMS

	err := decoder.Decode(&t)
	if err != nil {
		errRes(w, http.StatusBadRequest)
		return
	}

	err = sms.Send(t)
	if err != nil {
		errRes(w, http.StatusInternalServerError)
		return
	}

	successRes(w)
	return
}

func EmailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var t types.Email

	err := decoder.Decode(&t)
	if err != nil {
		errRes(w, http.StatusBadRequest)
		return
	}

	err = email.Send(t)
	if err != nil {
		errRes(w, http.StatusInternalServerError)
		return
	}

	successRes(w)
	return
}

func errRes(w http.ResponseWriter, code int) {
	resData := res{
		Success: false,
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resData)
}

func successRes(w http.ResponseWriter) {
	resData := res{
		Success: true,
	}

	json.NewEncoder(w).Encode(resData)
}

type res struct {
	Success bool `json:"success"`
}
