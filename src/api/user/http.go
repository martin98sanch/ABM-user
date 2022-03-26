package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/martin98sanch/ABM-user/src/api/server/response"
)

func Create(ctx *gin.Context) {
	response.Make(ctx, http.StatusOK, "")
}
