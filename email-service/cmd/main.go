package main

import (
	"fmt"
	"os"

	"github.com/MunizMat/microservices/email-service/internal/clients"
	"github.com/gin-gonic/gin"
)

func main() {
	clients.Init()

	port := os.Getenv("PORT")

	fmt.Println("Server running at http://localhost:", port, " 🚀")

	router := gin.Default()

	router.Run(fmt.Sprintf(":%s", port))
}
