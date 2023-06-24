// 描述在 GIN 中，如何自定义中间件

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

// 加入我们自定义一个统计各个路径请求次数的中间件
var mutex sync.Mutex
var path2Count = map[string]int64{}
func MetricCountOfPath() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		path2Count[ctx.Request.URL.Path] += 1
	}
}

func main() {
	app := gin.New()

	app.Use(MetricCountOfPath())
	app.GET("/count_of_path", func(ctx *gin.Context){
		mutex.Lock()

		respBody := make(gin.H, len(path2Count))
		for path, count := range path2Count {
			respBody[path] = count
		}
		mutex.Unlock()

		ctx.JSON(http.StatusOK, respBody)
	})

	app.Run()

}