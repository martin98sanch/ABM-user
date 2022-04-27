package request

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrCantReadBody = errors.New("can't read body")

func GetJsonBody(ctx *gin.Context, bodyStruct interface{}) error {
	err := json.NewDecoder(ctx.Request.Body).Decode(&bodyStruct)
	if err != nil {
		return ErrCantReadBody
	}
	return nil
}
