package services

import (
	"github.com/MunizMat/microservices/email-service/internal/models"
	"github.com/MunizMat/microservices/email-service/internal/repositories"
)

func SendEmail(email *models.Email) error {
	err := repositories.SaveEmail(email)

	return err
}
