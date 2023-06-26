// 介绍如何 redirect

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.New()

	app.GET("/baidu", func(ctx *gin.Context){
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})
	app.Run()
}