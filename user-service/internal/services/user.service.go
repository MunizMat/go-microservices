package services

import (
	"github.com/MunizMat/microservices/user-service/internal/models"
	"github.com/MunizMat/microservices/user-service/internal/repositories"
)

func CreateUser(user *models.User) error {
	err := repositories.CreateUser(user)

	return err
}
