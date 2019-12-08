package server

import (
	//jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {

	// It is good practice to version your API from the start
	v1 := router.Group("/api/v1")

	v1.GET("/", hello)
	v1.GET("/new", newfunction)
}
