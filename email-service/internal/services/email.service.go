package services

import (
	"fmt"
	"net/smtp"

	"github.com/MunizMat/microservices/email-service/internal/models"
	"github.com/MunizMat/microservices/email-service/internal/utils"
)

func SendEmail(user *models.User) {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	sender := utils.Environment.SMTP_SENDER_EMAIL
	password := utils.Environment.SMTP_SENDER_PASSWORD
	receiver := user.Email

	subject := "User Successfully Created\n"
	body := fmt.Sprintf("Congratulations on creating your user %s!\n", user.FirstName)
	message := []byte(subject + "\n" + body)

	auth := smtp.PlainAuth("", sender, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, sender, []string{receiver}, message)

	if err != nil {
		fmt.Println("Error sending email:", err)

	}

}
