package main

import (
	"fmt"
	"go-smtp/service"
	"log"
	"text/template"
)

func main() {
	m := service.Mail{
		FromEmailAdddress: "SENDER_EMAIL",
		FromName:          "SENDER_NAME",
		FromPassword:      "SENDER_PASSWORD",
		ToEmailAddress:    "RECIPIENT_EMAIL",
		Subject:           "Test Sendmail service by golang",
	}

	t, err := bodyTemplate(m, "template.html")

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%s", t.Body.Bytes())

	err = service.Mailer(*t)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Mail sent to: %s, subject: %s\n", m.ToEmailAddress, m.Subject)

}

func bodyTemplate(m service.Mail, filename string) (*service.Mail, error) {

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\nFrom: " + m.FromName + " <" + m.FromEmailAdddress + ">" + "\n"

	m.Body.Write([]byte(
		fmt.Sprintf("Subject: %s\n%s\n\n", m.Subject, mimeHeaders)))

	template, err := template.ParseFiles(filename)
	if err != nil {
		return nil, err
	}

	template.Execute(&m.Body, struct {
		Message string
		Name    string
		Body    string
		Email   string
	}{
		Message: "This is string message",
		Name:    m.FromName,
		Body:    "<a href=\"http://www.google.co.th\">click</a>",
		Email:   m.FromEmailAdddress,
	})
	return &m, nil

}
