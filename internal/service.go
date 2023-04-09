package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	gomail "gopkg.in/mail.v2"
)

type service struct {
	db       *mongo.Collection
	password string
	email    string
}

func NewService(db *mongo.Collection, p string, e string) *service {
	return &service{db: db, password: p, email: e}
}

func (s *service) SendEmail(ctx context.Context, email string, emailType string) {

	m := gomail.NewMessage()
	m.SetHeader("From", s.email)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "test subject")
	m.SetBody("text/plain", "This is Gomail test body")
	d := gomail.NewDialer("smtp.gmail.com", 587, s.email, s.password)

	//// This is only needed when SSL/TLS certificate is not valid on server.
	//// In production this should be set to false.
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	}

	//from := s.email
	//password := s.password
	//
	//// Receiver email address.
	//to := []string{
	//	email,
	//}
	//
	//// smtp server configuration.
	//smtpHost := "smtp.gmail.com"
	//smtpPort := "587"
	//
	//// Message.
	//message := []byte("This is a test email message.")
	//
	//// Authentication.
	//auth := smtp.PlainAuth("", from, password, smtpHost)
	//
	//// Sending email.
	//err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//log.Printf("Email to %s sent successfully\n", to)
	//
	//m := bson.M{"from": from, "to": to[0], "emailType": emailType}
	//one, err := s.db.InsertOne(ctx, m)
	//if err != nil {
	//	log.Printf("error due insert to database %v", err)
	//}
	//log.Printf("insrted to database %s", one.InsertedID)
}
