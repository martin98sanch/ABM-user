package server

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.Default()
	urlMapping(router)

	return router
}
