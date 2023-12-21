package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GIN 使用参考：https://gin-gonic.com/zh-cn/docs/examples/bind-body-into-dirrerent-structs/

func main() {
	app := gin.New()
	app.GET("/:key", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"key": ctx.Param("key"),
		})
	})

	app.GET("/:key1/name", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"key1": ctx.Param("key1"),
		})
	})

	app.Run()
}
