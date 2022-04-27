package lib

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(email string, token string) {

	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PW")

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	address := smtpHost + ":" + smtpPort

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	subject := "Subject: Password forgotted\r\n"
	body := "Hi! Your token is -> " + token
	message := []byte(subject + body)

	// Sending email.
	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}

}
//feMG9w==