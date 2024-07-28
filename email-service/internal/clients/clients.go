package clients

import (
	"log"

	"github.com/MunizMat/microservices/email-service/internal/utils"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type RabbitMQClient struct {
	Channel *amqp091.Channel
}

var (
	Redis    *redis.Client
	RabbitMQ *RabbitMQClient
)

func createRedisClient() {
	options, err := redis.ParseURL(utils.Environment.REDIS_URL)

	if err != nil {
		log.Fatal(err.Error())
	}

	Redis = redis.NewClient(options)
}

func recieveQueueMessage(messages <-chan amqp091.Delivery) {
	for message := range messages {
		log.Printf("Received a message: %s", message.Body)
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
	createRedisClient()
	createRabbitMQClient()
}
