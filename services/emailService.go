package services

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"example/BtcApp/repositories"
	"example/BtcApp/utils"
)

const (
	subjectTemplate     = "BTC to UAH rate"
	htmlContentTemplate = "<strong>Current rate: 1BTC = %f</strong>"
)

func AddEmail(email string) error {
	return repositories.AppendEmailToFile(email)
}

func SendRateForSubscribedEmails() error {
	emails, err := repositories.GetAllEmails()
	if err != nil {
		return err
	}

	exchangeRate, err := FetchExchangeRate()
	if err != nil {
		return err
	}

	from := mail.NewEmail(utils.Nickname, utils.EmailSender)
	htmlContent := fmt.Sprintf(htmlContentTemplate, exchangeRate.Value)

	client := sendgrid.NewSendClient(os.Getenv(utils.ApiKey))

	for _, email := range emails {
		to := mail.NewEmail("", email)
		plainTextContent := fmt.Sprintf("RATE: %f", exchangeRate.Value)
		message := mail.NewSingleEmail(from, subjectTemplate, to, plainTextContent, htmlContent)

		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
			return err
		}

		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return nil
}
