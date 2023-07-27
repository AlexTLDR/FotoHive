// experimenting with context

package main

import (
	"fmt"
	"os"

	"github.com/go-mail/mail/v2"
)

const (
	// was lazy to use a .env file but this is a sandbox account and I have regenerated the credentials
	host     = "sandbox.smtp.mailtrap.io"
	port     = 2525
	username = "ef98e50d183d83"
	password = "ecc7293fbc4307"
)

func main() {
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

	dialer := mail.NewDialer(host, port, username, password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Email sent!")
}
