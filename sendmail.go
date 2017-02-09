/*
 * HOW TO BUILD:
 * > go get gopkg.in/gomail.v2
 * > go get github.com/BurntSushi/toml
 * > go build
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"gopkg.in/gomail.v2"
)

func isEmpty(str string) bool {
	//return str == ""
	return len(str) == 0
}

// Config : app conf
type Config struct {
	sender    string
	recipient string
	server    string
	port      int
	subject   string
	body      string
	user      string
	password  string
}

func main() {

	// default parameters
	var conf = Config{
		sender:    "sender@example.com",
		recipient: "recipient@example.com",
		server:    "smtp.example.com",
		port:      25,
		subject:   "Email subject",
		body:      "Email body",
		user:      "",
		password:  "",
	}

	readConfFile(&conf)

	parseArguments(&conf)

	fmt.Printf("%+v\n", conf)
	fmt.Printf("CONNECTING TO: %s:%d\n", conf.server, conf.port)
	fmt.Printf("FROM: %s\n", conf.sender)
	fmt.Printf("TO: %s\n", conf.recipient)
	fmt.Printf("SUBJECT: %s\n", conf.subject)
	fmt.Printf("BODY: %s\n", conf.body)

	// authentication send
	if !isEmpty(conf.user) {
		fmt.Printf("user: %s\n", conf.user)
		if !isEmpty(conf.password) {
			fmt.Printf("password: ***\n")
		}
		fmt.Printf("\n\n")
		sendAuth(conf.server, conf.port, conf.sender, conf.recipient, conf.subject, conf.body, conf.user, conf.password)
		return
	}

	// without authentication
	fmt.Printf("\n\n")
	send(conf.server, conf.port, conf.sender, conf.recipient, conf.subject, conf.body)
	return
}

func readConfFile(conf *Config) {
	// read conf file
	tomlData, err := ioutil.ReadFile("sendmail.toml")
	if err != nil {
		fmt.Printf("error reading configuration file:\n\t%s\n", err)
	} else {
		// se il file è stato letto
		_, err := toml.Decode(string(tomlData), &conf)
		if err != nil {
			panic(err)
		}
	}
}

func parseArguments(conf *Config) {
	conf.sender = *flag.String("from", conf.sender, "sender email address")
	conf.recipient = *flag.String("to", conf.recipient, "recipient email address")
	conf.server = *flag.String("server", conf.server, "smtp server address")
	conf.port = *flag.Int("port", conf.port, "smtp server port")
	conf.body = *flag.String("body", conf.body, "email body")
	conf.subject = *flag.String("subject", conf.subject, "email subject")
	conf.user = *flag.String("user", conf.user, "authentication user")
	conf.password = *flag.String("password", conf.password, "authentication password")
	flag.Parse()
}

func send(smtpServerAddress string, smtpServerPort int, sender string, recipient string, subject string, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.Dialer{Host: smtpServerAddress, Port: smtpServerPort}
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		fmt.Println("")
		fmt.Println("")
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func sendAuth(smtpServerAddress string, smtpServerPort int, sender string, recipient string, subject string, body string, user string, password string) {
	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(smtpServerAddress, smtpServerPort, user, password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		fmt.Println("")
		fmt.Println("")
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
