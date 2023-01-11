package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

// Encapsulation
type SendEmailVerificationCodeParam struct {
	SenderEmail        string
	SendInBlueEmail    string
	SendInBluePassword string
	SendInBlueAddress  string
	SMTPHost           string
	UserName           string
	UserEmail          string
	OTP                string
}

func SendEmailVerificationCode(param SendEmailVerificationCodeParam) error {

	// Receiver email address.
	to := []string{
		param.UserEmail,
	}

	// Message.
	subject := "Email Verification Code"
	body := fmt.Sprintf("Hi %s ! \r\n"+
		"Your verification code is %s . \r\n"+
		"Enter this code in our Sanyukta APP to activate your account.\r\n"+
		"If you have any questions, send us an email %s.\r\n"+
		"We’re glad you’re here!\r\n"+
		"The Sanyukta team\r\n", param.UserName, param.OTP, param.SenderEmail)

	request := Mail{
		Sender:  param.SenderEmail,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	msg := BuildMessage(request)
	auth := smtp.PlainAuth("", param.SendInBlueEmail, param.SendInBluePassword, param.SMTPHost)
	err := smtp.SendMail(fmt.Sprintf("%s:587", param.SMTPHost), auth, param.SenderEmail, to, []byte(msg))
	if err != nil {
		logrus.Error("error while try to send otp email", err)
		return err
	}

	return err
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
