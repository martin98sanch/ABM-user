package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/martin98sanch/ABM-user/src/api/server/url"
)

func urlMapping(router *gin.Engine) {
	router.GET(url.Ping, ping)
	//router.POST(url.CreateUser, user.Create)

	router.NoRoute(endpointNotFound)
}

//ping returns a status 200 if the app is OK
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// endpointNotFound returns a status 404 if not exist the url
func endpointNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status_code": http.StatusNotFound,
		"message":     "Resource not found",
	})
}
