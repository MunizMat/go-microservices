package clients

import (
	"log"

	"github.com/MunizMat/microservices/user-service/internal/utils"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type RabbitMQClient struct {
	Channel           *amqp091.Channel
	UserCreationQueue *amqp091.Queue
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

	RabbitMQ = &RabbitMQClient{
		Channel:           channel,
		UserCreationQueue: &userCreationQueue,
	}

}

func Init() {
	createRedisClient()
	createRabbitMQClient()
}
