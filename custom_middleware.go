// 描述在 GIN 中，如何自定义中间件

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

// 加入一个统计各个路径请求次数的中间件
var mutex sync.Mutex
var path2Count = map[string]int64{}
func MetricCountOfPath() gin.HandlerFunc {
	// 这里还没有使用 c.Next()，如果要在调用下一个 handler 前后分别执行一段逻辑，则可以使用 c.Next()
	return func(ctx *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		path2Count[ctx.Request.URL.Path] += 1
	}
}

// MetricLatencyOfPath 一个统计请求处理时延的中间件
func MetricLatencyOfPath() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		// 调用下一个 中间件 / handler
		ctx.Next()
		latency := time.Now().Sub(start) / time.Millisecond
		fmt.Printf("Path: %s, latency: %d ms\n", ctx.Request.URL.Path, latency)
	}
}

func main() {
	app := gin.New()

	app.Use(MetricCountOfPath(), MetricLatencyOfPath())
	app.GET("/count_of_path", func(ctx *gin.Context){
		mutex.Lock()
		respBody := make(gin.H, len(path2Count))
		for path, count := range path2Count {
			respBody[path] = count
		}
		mutex.Unlock()
		time.Sleep(10 * time.Millisecond)
		ctx.JSON(http.StatusOK, respBody)
	})

	app.Run()

}