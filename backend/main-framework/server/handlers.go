package server

import (
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.String(200, "Hello World")
}

func newfunction(c *gin.Context) {
	c.String(200, "Hello changes")
}
