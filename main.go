package main

import (
	"io"
	"log"
	"net/mail"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("mail.msg")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	f.Close()

	log.Printf("email body read: %d bytes", len(body))

	r := strings.NewReader(string(body))

	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("message was read")

	header := m.Header
	log.Printf("Date: %s", header.Get("Date"))
	log.Printf("From: %s", header.Get("From"))
	log.Printf("To: %s", header.Get("To"))
	log.Printf("Subject: %s", header.Get("Subject"))

	body, err = io.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Body: %s", body)
}
