package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	validHTTPCodes = []int{
		http.StatusOK,
		http.StatusCreated,
		http.StatusNoContent,
	}
)

//TODO: Anda medio mal, en los status codes negativos, no postea el body en message
// En los success, agarra todo en un array
func Make(ctx *gin.Context, statusCode int, body ...interface{}) {
	var messageBody interface{}
	messageBody = gin.H{
		"status_code": statusCode,
		"message":     body,
	}

	for _, validStatusCode := range validHTTPCodes {
		if statusCode == validStatusCode {
			messageBody = body
		}
	}

	ctx.JSON(statusCode, messageBody)
}
