package repositories

import (
	"context"
	"fmt"

	"github.com/MunizMat/microservices/user-service/internal/clients"
	"github.com/MunizMat/microservices/user-service/internal/models"
)

func CreateUser(user *models.User) error {
	key := fmt.Sprintf("user:%s", user.Id)

	status := clients.Redis.JSONSet(context.Background(), key, "$", *user)

	_, err := status.Result()

	return err
}
