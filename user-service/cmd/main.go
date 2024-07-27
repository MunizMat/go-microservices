package main

import (
	"fmt"
	"os"

	"github.com/MunizMat/api/internal/clients"
	"github.com/MunizMat/api/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
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
