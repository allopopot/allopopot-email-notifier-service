package queues

import (
	"allopopot-email-service/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"gopkg.in/gomail.v2"
)

type Attachments struct {
	Filename string `json:"filename"`
	MimeType string `json:"mimetype"`
	Payload  string `json:"payload"`
}

type EmailPayload struct {
	To          []string      `json:"to"`
	Subject     string        `json:"subject"`
	Body        string        `json:"body"`
	Attachments []Attachments `json:"attachments"`
}

func (at *Attachments) WriteToFile() string {
	uid, err := uuid.NewV7()
	if err != nil {
		log.Panic(err.Error())
	}
	tempPath := fmt.Sprintf("./temp/%s", uid.String())
	// extension, _ := mime.ExtensionsByType(at.MimeType)
	tempPathFile := fmt.Sprintf("%s/%s", tempPath, at.Filename)

	os.MkdirAll(tempPath, os.FileMode(os.O_CREATE))
	dataBytes, _ := base64.RawStdEncoding.DecodeString(at.Payload)

	file, _ := os.Create(tempPathFile)
	defer file.Close()
	file.Write(dataBytes)
	file.Sync()

	return tempPathFile
}

func SendMail(payload *EmailPayload) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.SMTP_SENDER)
	m.SetHeader("To", payload.To...)
	m.SetHeader("Subject", payload.Subject)
	m.SetBody("text/html", payload.Body)

	if len(payload.Attachments) > 0 {
		for _, v := range payload.Attachments {
			path := v.WriteToFile()
			m.Attach(path)
			defer os.RemoveAll(path)
		}
	}

	d := gomail.NewDialer(config.SMTP_HOST, config.SMTP_PORT, config.SMTP_USERNAME, config.SMTP_PASSWORD)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func InitEmailDispatcherQueue(channel *amqp.Channel) {

	const QueueName = "email-dispatch"

	q, err := channel.QueueDeclare(QueueName, true, false, true, false, nil)
	if err != nil {
		log.Panicln("Failed to declare queue.")
	}
	log.Println("Queue declared.")

	err = channel.QueueBind(q.Name, "", config.AMQP_EXCHANGE_NAME, false, nil)
	if err != nil {
		log.Panicln("Failed to bind queue to exchange.")
	}
	log.Println("Queue bind to exhange successful.")

	msgs, err := channel.Consume(QueueName, "", false, true, false, false, nil)
	if err != nil {
		log.Panicln("Failed to consume messages.", err)
	}
	log.Println("Ready to consume messages.")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for d := range msgs {
			payload := new(EmailPayload)
			json.Unmarshal(d.Body, payload)
			err := SendMail(payload)
			if err != nil {
				d.Nack(false, true)
			}
			d.Ack(false)
		}
	}()
	wg.Wait()

}
