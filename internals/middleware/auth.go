package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/models"
	"github.com/sai-mudike/careerdock.git/internals/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {

		context.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Code: customErr.ErrUnauthorized.Error(), Message: customErr.ErrUnauthorized.Error()})
		return
	}
	userID, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Code: err.Error(), Message: err.Error()})
		return
	}

	context.Set("userID", userID)
	context.Next()

}
