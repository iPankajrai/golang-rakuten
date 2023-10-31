package handlers

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
	//c.String(200, "pong")s
	c.Abort()
}
