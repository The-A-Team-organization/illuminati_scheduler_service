package main

import (
	"fmt"
	"os"

	"github.com/wneessen/go-mail"
)

func main() {
	username := ""
	password := ""
	to       := ""
	client, err := mail.NewClient("smtp.gmail.com", mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover), mail.WithUsername(username), mail.WithPassword(password))	
	if err != nil {
		fmt.Printf("failed to create mail client: %s\n", err)
		os.Exit(1)
	}

        msg := mail.NewMsg()
        msg.Subject("Word of the Day")

       body := `Hello,

You’re subscribed to our English vocabulary service, and today’s Word of the Day is corruption.

Learn its meaning, usage, and examples to expand your vocabulary!

Happy learning,
— The English Words Team
`
        msg.SetBodyString(mail.TypeTextPlain, body)

        msg.From(username)
	msg.To(to)
	if err = client.DialAndSend(msg); err != nil {
		fmt.Printf("failed to send mail: %s\n", err)
		os.Exit(1)
	}
}
