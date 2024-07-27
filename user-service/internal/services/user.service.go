package services

import (
	"github.com/MunizMat/api/internal/models"
	"github.com/MunizMat/api/internal/repositories"
)

func CreateUser(user *models.User) error {
	err := repositories.CreateUser(user)

	return err
}
