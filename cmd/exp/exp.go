// experimenting with context

package main

import (
	"os"

	"github.com/go-mail/mail/v2"
)

func main() {
	from := "alex@alex.com"
	to := "alex@tldr.com"
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
}
