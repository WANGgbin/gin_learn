// 描述如何优雅的关闭服务器

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := gin.New()

	// 采用一个协程启动 app
	go func() {
		if err := app.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server run error: %v", err)
		}
	}()

	// 注册监听信号
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGQUIT)
	<-sigCh

	// TODO(@wangguobin): 添加 engine.Serve() 方法, 返回底层的 Server()
	if err := app.Server().Shutdown(); err != nil {
		log.Fatalf("shutdown server, error: %v", err)
	}
	log.Printf("shutdown server gracefully!")
}