package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/sai-mudike/careerdock.git/internals/customErr"

	"github.com/gin-gonic/gin"
)

func HandleErrorWithGin(context *gin.Context, err error) {

	var custErr *customErr.CustomErr

	if errors.As(err, &custErr) {
		context.JSON(custErr.StatusCode, gin.H{
			"code":    custErr.Code,
			"message": custErr.Message,
		})

		return
	}

	log.Printf("internal error: %v", err)

	context.JSON(http.StatusInternalServerError, gin.H{
		"code":    "INTERNAL_SERVER_ERROR",
		"message": "something went wrong",
	})

}
