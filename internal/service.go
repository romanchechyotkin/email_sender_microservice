package service

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(email string) {

	from := "testcarbookingservice@gmail.com"
	password := "buttaxburfjwnerg"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("Email to %s sent successfully\n", to)
}
