package main

import (
	"github.com/MunizMat/microservices/email-service/internal/clients"
	"github.com/MunizMat/microservices/email-service/internal/utils"
)

func main() {
	utils.ParseEnvsOrPanic()

	clients.Init()

}
