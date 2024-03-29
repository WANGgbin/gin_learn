package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GIN 使用参考：https://gin-gonic.com/zh-cn/docs/examples/bind-body-into-dirrerent-structs/

func main() {

	app := gin.New()
	app.GET("/:usr/name", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"usr": ctx.Param("usr"),
		})
	})

	app.GET("/:student/age", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"student": ctx.Param("student"),
		})
	})

	app.Run()
}