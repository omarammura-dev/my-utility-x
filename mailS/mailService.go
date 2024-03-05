package mailS

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendSimpleMessage(url, userEmail, username, templateId string) (string, error) {
	from := mail.NewEmail("Omar Ammura", "no-reply@ammura.tech")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail(username, userEmail)
	contents := mail.NewContent("text/html", "l")
	message := mail.NewV3MailInit(from, subject, to, contents)

	message.Personalizations[0].SetDynamicTemplateData("name", username)
	message.Personalizations[0].SetDynamicTemplateData("callbackUrl", url)
	message.SetTemplateID(templateId)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {

		log.Println(err)
		return "", err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return response.Body, nil
}
