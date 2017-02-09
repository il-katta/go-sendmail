package main

import (
	"bytes"
	"flag"
	"log"
	"net/smtp"
	"strconv"
)

func main() {
	var sender = flag.String("s", "sender@example.com", "sender email address")
	var recipient = flag.String("r", "recipient@example.com", "recipient email address")
	var smtpServerAddress = flag.String("srv", "", "smtp server address")
	var smtpServerPort = flag.Int("p", 25, "smtp server port")

	flag.Parse()

	// Connect to the remote SMTP server.
	c, err := smtp.Dial(*smtpServerAddress + ":" + strconv.Itoa(*smtpServerPort))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	// Set the sender and recipient.
	c.Mail(*sender)
	c.Rcpt(*recipient)
	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	buf := bytes.NewBufferString("test email")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}
