package mail

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)


func SendSimpleMessage() (string, error) {
	from := mail.NewEmail("Omar Ammura", "no-reply@ammura.tech")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Alperen Alp", "alperenalp216@gmail.com")
	plainTextContent := "Merhaba bu test mesaji gorursen 100 bin TL kazandin"
	htmlContent := "<p>Merhaba bu test mesaji gorursen 100 bin TL kazandin</p>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		
		log.Println(err)
		return "",err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return response.Body,nil
}
