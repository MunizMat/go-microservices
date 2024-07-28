package main

import (
	"fmt"

	"github.com/MunizMat/microservices/user-service/internal/clients"
	"github.com/MunizMat/microservices/user-service/internal/routes"
	"github.com/MunizMat/microservices/user-service/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.ParseEnvsOrPanic()

	clients.Init()

	defer clients.RabbitMQ.Connection.Close()
	defer clients.RabbitMQ.Channel.Close()

	fmt.Println("Server up and running at http://localhost:", utils.Environment.PORT, " 🚀")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
	}))

	// User
	router.POST("/users", routes.UserPost)

	router.Run(fmt.Sprintf(":%s", utils.Environment.PORT))
}
