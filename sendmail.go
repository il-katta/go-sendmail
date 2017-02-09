/*
 * HOW TO BUILD:
 * > go get gopkg.in/gomail.v2
 * > go get gopkg.in/yaml.v2
 * > go build
 */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/gomail.v2"
	"gopkg.in/yaml.v2"
)

func isEmpty(str string) bool {
	//return str == ""
	return len(str) == 0
}

// Config : app conf
type Config struct {
	Sender    string `yaml:"sender"`
	Recipient string `yaml:"recipient"`
	Server    string `yaml:"server"`
	Port      int    `yaml:"port"`
	Subject   string `yaml:"subject"`
	Body      string `yaml:"body"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
}

func main() {

	// default parameters
	var conf = Config{
		Sender:    "sender@example.com",
		Recipient: "recipient@example.com",
		Server:    "smtp.example.com",
		Port:      25,
		Subject:   "Email subject",
		Body:      "Email body",
		User:      "",
		Password:  "",
	}

	readConfFile(&conf)

	sender := flag.String("from", conf.Sender, "sender email address")
	recipient := flag.String("to", conf.Recipient, "recipient email address")
	server := flag.String("server", conf.Server, "smtp server address")
	port := flag.Int("port", conf.Port, "smtp server port")
	body := flag.String("body", conf.Body, "email body")
	subject := flag.String("subject", conf.Subject, "email subject")
	user := flag.String("user", conf.User, "authentication user")
	password := flag.String("password", conf.Password, "authentication password")
	flag.Parse()

	conf.Sender = *sender
	conf.Recipient = *recipient
	conf.Server = *server
	conf.Port = *port
	conf.Body = *body
	conf.Subject = *subject
	conf.User = *user
	conf.Password = *password

	fmt.Printf("CONNECTING TO: %s:%d\n", conf.Server, conf.Port)
	fmt.Printf("FROM: %s\n", conf.Sender)
	fmt.Printf("TO: %s\n", conf.Recipient)
	fmt.Printf("SUBJECT: %s\n", conf.Subject)
	fmt.Printf("BODY: %s\n", conf.Body)
	// authentication send
	if !isEmpty(conf.User) {
		fmt.Printf("user: %s\n", conf.User)
		if !isEmpty(conf.Password) {
			fmt.Printf("password: ***\n")
		}
	}

	// without authentication
	fmt.Printf("\n\n")
	send(conf)
}

func readConfFile(conf *Config) {
	// read conf file
	yamlData, err := ioutil.ReadFile("sendmail.yml")
	if err != nil {
		fmt.Printf("error reading configuration file:\n\t%s\n", err)
	} else {
		// se il file è stato letto
		err := yaml.Unmarshal(yamlData, &conf)
		fmt.Printf("\n\nconf file %v\n\n", conf.Sender)
		if err != nil {
			panic(err)
		}
	}
}

func send(conf Config) {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.Sender)
	m.SetHeader("To", conf.Recipient)
	m.SetHeader("Subject", conf.Subject)
	m.SetBody("text/plain", conf.Body)

	var d *gomail.Dialer
	if isEmpty(conf.User) {
		d = &gomail.Dialer{Host: conf.Server, Port: conf.Port}
	} else {
		d = gomail.NewDialer(conf.Server, conf.Port, conf.User, conf.Password)
	}
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		fmt.Println("")
		fmt.Println("")
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
