# Microservices in Practice With Go 

This is a simple implementation of a user creation flow that enables two applications to communicate with each other through a message broker, which, in this case, is RabbitMQ. This repository consists of two applications / services, a "user" service and a "email" service. Each service is a Go application that runs indpendently.

The user service acts as a producer that publishes messages to a RabbitMQ queue whenever a user is created. The queue messages are then consumed by the email service.

This is a very simple implementation and I am aware that there are easier ways to perform this. However, my focus here, besides familiarizing myself with Go, was to better understand microservice architecture and messaging systems through a real implementation.

## Technologies
- Redis
- Go
- RabbitMQ
- Docker

## Running Locally
Prerequisites:
- Docker
- An email + password that can be used for SMTP authentication

Start by cloning the repository:

```
git clone https://github.com/MunizMat/go-microservices.git
```

Then setup the necessary .env files by running the following commands in your terminal:
```
echo "RABBITMQ_DEFAULT_USER=\nRABBITMQ_DEFAULT_PASS=" > .env

cd user-service && echo "PORT=\nREDIS_URL=\nRABBIT_MQ_URL=" > .env

cd ../email-service && echo "SMTP_SENDER_EMAIL=\SMTP_SENDER_PASSWORD=\nRABBIT_MQ_URL=" > .env
```

Use docker compose to run the applications:

```
docker compose up --build
```

You can also use the .dev compose file, which uses [Air](https://github.com/air-verse/air) to run the apps, enabling live reloading:
```
docker compose -f compose.dev.yml up --build
```
