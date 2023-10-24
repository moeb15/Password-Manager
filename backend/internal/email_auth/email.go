package email

import (
	"fmt"
	"os"
	"pwdmanager_api/pkg/models"

	"gopkg.in/gomail.v2"
)

func RegistrationEmail(user models.User) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("OWNER_EMAIL"))
	msg.SetHeader("To", user.Email)
	msg.SetHeader("Subject", "Account Created")
	msg.SetBody("text/html", fmt.Sprintf("Account creattion successful <b>%s</b>", user.Username))

	dialer := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("OWNER_EMAIL"),
		os.Getenv("OWNER_PWD"))

	err := dialer.DialAndSend(msg)
	return err
}
