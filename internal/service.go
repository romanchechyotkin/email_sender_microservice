package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/smtp"
)

type service struct {
	db *mongo.Collection
}

func NewService(db *mongo.Collection) *service {
	return &service{db: db}
}

func (s *service) SendEmail(ctx context.Context, email string, emailType string) {

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

	m := bson.M{"from": from, "to": to[0], "emailType": emailType}
	one, err := s.db.InsertOne(ctx, m)
	if err != nil {
		log.Printf("error due insert to database %v", err)
	}
	log.Printf("insrted to database %s", one.InsertedID)
}
