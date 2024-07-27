package repositories

import (
	"context"
	"fmt"

	"github.com/MunizMat/microservices/email-service/internal/clients"
	"github.com/MunizMat/microservices/email-service/internal/models"
)

func SaveEmail(email *models.Email) error {
	key := fmt.Sprintf("email:%s", email.Id)

	status := clients.Redis.JSONSet(context.Background(), key, "$", *email)

	_, err := status.Result()

	return err
}
