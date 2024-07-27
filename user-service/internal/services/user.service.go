package services

import (
	"encoding/json"
	"fmt"

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

	clients.RabbitMQ.Channel.Publish("", clients.RabbitMQ.UserCreationQueue.Name, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        userJson,
	})

	return err
}
