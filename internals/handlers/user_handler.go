package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/services"
)

func RegisterUser(context *gin.Context) {

	ctx := context.Request.Context()

	var newUser models.User

	err := context.ShouldBindJSON(&newUser)

	if err != nil {
		HandleError(context, customErr.ErrInvalidRequest)
		return
	}

	err = services.CreateUser(ctx, newUser)

	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})

}

func UserLogin(context *gin.Context) {

	ctx := context.Request.Context()

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		HandleError(context, customErr.ErrInvalidRequest)
		return
	}

	token, err := services.UserLogin(ctx, user)

	if err != nil {
		HandleError(context, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "loged in successfully", "token": token})

}
