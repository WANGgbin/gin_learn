package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	app := gin.New()

	app.GET("/cookie", func(ctx *gin.Context){
		// 本质上调用 http.Request.Cookie() 方法
		cookie, err := ctx.Cookie("gin-cookie")

		// 还没有为客户端设置 cookie
		if err == http.ErrNoCookie {
			cookie = "not-set"
			ctx.SetCookie("gin-cookie", "text", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("cookie: %s\n", cookie)
	})
	app.Run()
}