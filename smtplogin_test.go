package smtplogin

import (
	"testing"
	"net/smtp"
	"log"
)

func TestLoginAuth(t *testing.T) {
	auth := LoginAuth("", "username", "P@ssw0rd", "smpp.server.com")

	// Connect to the remote SMTP server.
	c, err := smtp.Dial("smpp.server.com:25")
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Auth(auth); err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
	if err := c.Mail("to@mail.com"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("from@mail.com"); err != nil {
		log.Fatal(err)
	}


	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}

	_, err = wc.Write([]byte("This is the email body"))

	if err != nil {
		log.Fatal(err)
	}
	wc.Close()
	c.Quit()

}
