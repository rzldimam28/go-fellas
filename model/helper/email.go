package helper

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/gomail.v2"
)

func SendMailVerivication(toEmailAddress string, userId primitive.ObjectID) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "imamrizaldi00@gmail.com")
	msg.SetHeader("To", toEmailAddress)
	msg.SetHeader("Subject", "Email Verifications")
	userIdString := userId.Hex()
	msg.SetBody("text/html", "Klik Link Berikut Untuk Verifikasi: go-fellas.herokuapp.com/users/verify/" + userIdString)

	n := gomail.NewDialer("smtp.gmail.com", 587, "imamrizaldi00@gmail.com", "kahthvprcnhzkowh")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		return errors.New("can not send mail")
	}
	return nil
}

func SendMailCommentNotif(toEmailAddress string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "imamrizaldi00@gmail.com")
	msg.SetHeader("To", toEmailAddress)
	msg.SetHeader("Subject", "Comment Notifications")
	msg.SetBody("text/html", "Blog Anda Baru saja Dikomentari")

	n := gomail.NewDialer("smtp.gmail.com", 587, "imamrizaldi00@gmail.com", "kahthvprcnhzkowh")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}

func SendMailLikeNotif(toEmailAddress string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "imamrizaldi00@gmail.com")
	msg.SetHeader("To", toEmailAddress)
	msg.SetHeader("Subject", "Like Notifications")
	msg.SetBody("text/html", "Blog Anda Baru saja Dilike")

	n := gomail.NewDialer("smtp.gmail.com", 587, "imamrizaldi00@gmail.com", "kahthvprcnhzkowh")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		return err
	}
	return nil
}