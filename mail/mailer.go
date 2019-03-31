package mail

import (
	"context"
	"fmt"
	"os"
	"time"

	mailgun "github.com/mailgun/mailgun-go/v3"
)

type Mailer interface {
	SendEmail(to string, subject string, text string) error
}

func NewMailer() Mailer {
	return mailgunMailer{}
}

type mailgunMailer struct{}

func (s mailgunMailer) SendEmail(to string, subject string, text string) error {

	domain := os.Getenv("MAIL_DOMAIN")
	privateAPIKey := os.Getenv("MAILGUN_PRIVATE_API_KEY")
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(domain, privateAPIKey)

	sender := "no-reply@hubb-cloud.com"

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, text, to)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message	with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}
