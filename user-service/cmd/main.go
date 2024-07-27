package main

import (
	"fmt"
	"os"

	"github.com/MunizMat/microservices/user-service/internal/clients"
	"github.com/MunizMat/microservices/user-service/internal/routes"
	"github.com/MunizMat/microservices/user-service/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.ParseEnvsOrPanic()

	clients.Init()

	port := os.Getenv("PORT")

	fmt.Println("Server running at http://localhost:", port, " 🚀")

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
	}))

	// User
	router.POST("/users", routes.UserPost)

	router.Run(fmt.Sprintf(":%s", port))
}
