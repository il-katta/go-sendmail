/*
 * HOW TO BUILD:
 * > go get gopkg.in/gomail.v2
 * > go build
 */
package main

import (
    "flag"
    "gopkg.in/gomail.v2"
    "fmt"
    "bufio"
    "os"
)

func main() {
    var sender = flag.String("from", "sender@example.com", "sender email address")
    var recipient = flag.String("to", "recipient@example.com", "recipient email address")
    var smtp_server_address = flag.String("server", "aspmx.l.google.com", "smtp server address")
    var smtp_server_port = flag.Int("port", 25, "smtp server port")
    var body = flag.String("body", "Email test", "email body")
    var subject = flag.String("subject", "Email test", "email subject")
    flag.Parse()
    fmt.Printf("CONNECTING TO: %s:%d\n", *smtp_server_address, *smtp_server_port )
    fmt.Printf("FROM: %s\n", *sender)
    fmt.Printf("TO: %s\n", *recipient)
    fmt.Printf("SUBJECT: %s\n", *subject)
    fmt.Printf("BODY: %s\n", *body)
    fmt.Printf("\n\n")
	send(*smtp_server_address, *smtp_server_port, *sender, *recipient, *subject, *body)
}

func send(smtp_server_address string, smtp_server_port int, sender string, recipient string, subject string, body string) {
	m := gomail.NewMessage()
    m.SetHeader("From", sender)
    m.SetHeader("To", recipient)
    m.SetHeader("Subject", subject)
    m.SetBody("text/plain", body)

    d := gomail.Dialer{Host: smtp_server_address, Port: smtp_server_port}
    if err := d.DialAndSend(m); err != nil {
        fmt.Println(err)
        fmt.Println("")
        fmt.Println("")
    }

    fmt.Print("Press 'Enter' to continue...")
    bufio.NewReader(os.Stdin).ReadBytes('\n') 
}