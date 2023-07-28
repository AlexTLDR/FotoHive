// experimenting with context

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-mail/mail/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("SMTP_HOST")
	portString := os.Getenv("SMTP_PORT")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	from := "alex@alex.com"
	to := "alex.badragan@protonmail.com"
	subject := "hello"
	plaintext := "hello world from the email body"
	html := "<h1>Hello World</h1><p>hello world from the email body</p><p>more text</p>"

	msg := mail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	msg.WriteTo(os.Stdout)
	port, _ := strconv.Atoi(portString)
	dialer := mail.NewDialer(host, port, username, password)
	err = dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Email sent!")
}
