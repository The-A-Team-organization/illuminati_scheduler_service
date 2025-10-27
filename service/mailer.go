package service

import (
	"encoding/json"
	"errors"
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

func GetRandomWord() string {

	loremIpsumGenerator := loremipsum.New()
	Word := loremIpsumGenerator.Word()
	return strings.TrimSpace(Word)

}

func GetAppParticipants(url string) ([]string, error) {

	resp, err  := http.Get(url)
	if err != nil {
		fmt.Println("SOME TEXT:", err)
		return nil, err
	}

	defer resp.Body.Close()
	
	var data struct {
	 	participants []string `json:"participants"`
	}

	json.NewDecoder(resp.Body).Decode(&data);

	return data.participants, nil
}

func SendWordEmail(word string, participants []string) error{

	if(word == "") {
		return errors.New("the word field is blank")
	}

	client, err := mail.NewClient(
		"smtp.gmail.com",
		mail.WithTLSPortPolicy(mail.TLSMandatory),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(username),
		mail.WithPassword(password),
	)
	if err != nil {
		fmt.Println("Failed to create mail client:", err)
		return err
	}

	msg := mail.NewMsg()
	msg.Subject("Word of the Day")

	body := fmt.Sprintf(`Hello,

		You're subscribed to our Latin vocabulary service, and today’s Word of the Day is %s.

		Learn its meaning, usage, and examples to expand your vocabulary!

		Happy learning
		— The Latin Words Team
		`, 
	word)

	msg.SetBodyString(mail.TypeTextPlain, body)

	for _, email := range participants {
		msg.From(username)
		msg.To(email)

		if err := client.DialAndSend(msg); err != nil {
			fmt.Println("Failed to send mail:", err)
			return err
		}

		log.Println("Email sent successfully!")
	}

	return nil
}
