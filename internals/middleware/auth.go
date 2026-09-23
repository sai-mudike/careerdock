package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sai-mudike/careerdock.git/internals/customErr"
	"github.com/sai-mudike/careerdock.git/internals/handlers"
	"github.com/sai-mudike/careerdock.git/internals/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		handlers.HandleErrorWithGin(context, customErr.New(
			customErr.CodeUnauthorized,
			"authentication required",
			http.StatusUnauthorized,
			nil,
		))
		context.Abort()
		return
	}
	userID, err := utils.VerifyToken(token)
	if err != nil {
		handlers.HandleErrorWithGin(context, err)

		context.Abort()
		return
	}

	context.Set("userID", userID)
	context.Next()

}
