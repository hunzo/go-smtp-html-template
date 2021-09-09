package service

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
)

type Mail struct {
	FromEmailAdddress string
	FromName          string
	FromPassword      string
	ToEmailAddress    string
	Subject           string
	Body              bytes.Buffer
}

func Mailer(m Mail) error {
	smtpHost := "smtp.office365.com"

	conn, err := net.Dial("tcp", "smtp.office365.com:587")
	if err != nil {
		log.Fatal(err)
	}

	c, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Fatal(err)
	}

	tlsconfig := &tls.Config{
		ServerName: smtpHost,
	}

	if err := c.StartTLS(tlsconfig); err != nil {
		log.Fatal(err)
	}

	auth := LoginAuth(m.FromEmailAdddress, m.FromPassword)

	if err = c.Auth(auth); err != nil {
		println(err)
	}

	err = smtp.SendMail("smtp.office365.com:587", auth, m.FromEmailAdddress, []string{m.ToEmailAddress}, m.Body.Bytes())
	if err != nil {
		fmt.Println(err)
	}
	return nil

}
