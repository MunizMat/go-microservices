package routes

import (
	"fmt"
	"net/http"

	"github.com/MunizMat/microservices/user-service/internal/models"
	"github.com/MunizMat/microservices/user-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func UserPost(context *gin.Context) {
	var user models.User
	validate := validator.New()

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.CreateUser(&user)

	if err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Created user sucessfully", "user": user})
}
