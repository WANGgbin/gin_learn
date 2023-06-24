package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = map[string]gin.H {
	"wgb": {
		"name": "wgb",
		"email": "xxx@qq.com",
	},
}

func main() {
	app := gin.New()

	// 创建一个 routeGroup
	authGroup := app.Group("/auth",
		// 关于 http 的基本验证，可以参考：https://zhuanlan.zhihu.com/p/64584734
		// 服务端对 req 的 Header "Authorization" 进行校验，如果校验不通过返回 401 同时
		// 设置 Header：www-Authenticate: Basic realm= "family"
		// Basic: 表示需要基本认证。Basic 的内容为 username:password 的 base64 格式
		gin.BasicAuth(
			gin.Accounts{
				"wgb": "12345",
			},
		),
	)

	// 往 routeGroup 中注册路由
	authGroup.GET("get_secret", func(ctx *gin.Context){
		userName := ctx.MustGet(gin.AuthUserKey).(string)
		if secret, exist := secrets[userName]; exist {
			ctx.JSON(http.StatusOK, secret)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"reason": "no secret",
			})
		}

	})

	// 启动服务
	app.Run()
}