package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quentinguidee/microservice-core/middleware"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ErrorMiddleware())

	router.GET("/ping", ping)

	return router
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
