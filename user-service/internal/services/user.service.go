package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/MunizMat/microservices/user-service/internal/clients"
	"github.com/MunizMat/microservices/user-service/internal/models"
	"github.com/MunizMat/microservices/user-service/internal/repositories"
	"github.com/rabbitmq/amqp091-go"
)

func CreateUser(user *models.User) error {
	err := repositories.CreateUser(user)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	userJson, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = clients.RabbitMQ.Channel.PublishWithContext(ctx, "", clients.RabbitMQ.UserCreationQueue.Name, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        userJson,
	})

	return err
}
