// 介绍 GIN 中如何将不同格式的 req.body 解析到结构体中

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type PersonInfo struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
}

func main() {
	app := gin.New()

	app.POST("bind_req", func(ctx *gin.Context){
		// gin 提供了一系列 Bind 方法 用来解析 req.Body
		personInfo := &PersonInfo{}
		err := ctx.BindJSON(personInfo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "body should be json",
			})
			return
		}

		log.Printf("personInfo: %+v", personInfo)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	app.Run()
}