package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/:name/:gender", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			c.Params[0].Key: c.Params[0].Value,
			c.Params[1].Key: c.Params[1].Value,
		})
	})
	r.Run() 
}