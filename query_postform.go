// query && post form

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.New()

	app.GET("post", func(ctx *gin.Context){
		// Query/PostForm 底层都是调用 url.ParseQuery() 方法
		id := ctx.Query("id")
		name := ctx.PostForm("name")

		fmt.Printf("id: %s, name: %s", id, name)
		ctx.String(http.StatusOK, "success")
	})
	app.Run()
	gin.LoggerWithConfig()
}