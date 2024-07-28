package clients

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/MunizMat/microservices/email-service/internal/models"
	"github.com/MunizMat/microservices/email-service/internal/services"
	"github.com/MunizMat/microservices/email-service/internal/utils"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	Channel *amqp091.Channel
}

var (
	RabbitMQ *RabbitMQClient
)

func recieveQueueMessage(messages <-chan amqp091.Delivery) {
	for message := range messages {
		var user models.User

		err := json.Unmarshal(message.Body, &user)

		if err != nil {
			fmt.Println("Error receiving message: ", err.Error())
			continue
		}

		services.SendEmail(&user)
	}
}

func createRabbitMQClient() {

	connection, err := amqp091.Dial(utils.Environment.RABBIT_MQ_URL)

	utils.FailOnError(err, "Failed to connect to RabbitMQ")

	defer connection.Close()

	channel, err := connection.Channel()

	utils.FailOnError(err, "Failed to open RabbitMQ channel")

	defer channel.Close()

	userCreationQueue, err := channel.QueueDeclare(
		"user-creation",
		false,
		false,
		false,
		false,
		nil,
	)

	utils.FailOnError(err, "Failed to declare a queue")

	messages, err := channel.Consume(userCreationQueue.Name, "", true, false, false, false, nil)

	utils.FailOnError(err, "Failed to register RabbitMQ consumer")

	rabbitMqChan := make(chan int)

	go recieveQueueMessage(messages)

	RabbitMQ = &RabbitMQClient{
		Channel: channel,
	}

	log.Println("Listening for messages...")

	<-rabbitMqChan
}

func Init() {

	createRabbitMQClient()
}
