package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vertex-center/vertex-core-golang/router/middleware"
)

// CreateRouter creates a new Gin Router, using cors configuration passed as parameter.
// It returns the main router, and an API route group.
func CreateRouter(cors gin.HandlerFunc) (*gin.Engine, *gin.RouterGroup) {
	router := gin.Default()
	router.Use(cors)
	router.Use(middleware.ErrorMiddleware())

	router.GET("/ping", ping)

	api := router.Group("/api")
	api.GET("/ping", ping)

	return router, api
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
