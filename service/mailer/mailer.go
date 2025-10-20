package mailer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/XANi/loremipsum"
	"github.com/wneessen/go-mail"
)

var (
	 username = os.Getenv("EMAIL_USERNAME") 
	 password = os.Getenv("EMAIL_PASSWORD")
     mockURL  = os.Getenv("MOCK_URL")
)

func getRandomWord() string {
	loremIpsumGenerator := loremipsum.New()
	Word := loremIpsumGenerator.Word()
	return strings.TrimSpace(Word)

}

func SendWordEmail() {
	word := getRandomWord() 
	
	resp, err  := http.Get(mockURL)
	defer resp.Body.Close()
	
	var data struct {
		Participants []string `json:"participants"`
	}
	json.NewDecoder(resp.Body).Decode(&data);


	client, err := mail.NewClient(
		"smtp.gmail.com",
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(username),
		mail.WithPassword(password),
	)
	if err != nil {
		fmt.Println("Failed to create mail client:", err)
		return
	}

	msg := mail.NewMsg()
	msg.Subject("Word of the Day")

	body := fmt.Sprintf(`Hello,

You're subscribed to our Latin vocabulary service, and today’s Word of the Day is %s.

Learn its meaning, usage, and examples to expand your vocabulary!

Happy learning
— The Latin Words Team
`, word)

	msg.SetBodyString(mail.TypeTextPlain, body)

for _, email := range data.Participants {
	msg.From(username)
	msg.To(email)

	if err := client.DialAndSend(msg); err != nil {
		fmt.Println("Failed to send mail:", err)
		return
	}

	log.Println("Email sent successfully!")
}
}

