package email

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

type Email struct {
	Emails  []string `json:"emails"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func NewEmail() *Email {
	return &Email{}
}

type MainSender struct {
	From   string
	Dailer *gomail.Dialer
}

func NewMainSender() *MainSender {
	return &MainSender{}
}

func (ms *MainSender) Send(emailChan chan Email) error {
	m := gomail.newMessage()
	m.SetHeader("From", ms.From)
	for ec := range emailChan {
		m.SetHeader("Subject", ec.Subject)
		m.SetBody("text/html", ec.Body)
		for _, to := range ec.Emails {
			m.SetHeader("To", to)
			if err := ms.Dialer.DialAndSend(m); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
