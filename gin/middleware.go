package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func usedTime() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		// 需要再next之前设置
		context.Set("user", "irene")
		context.Next()
		used := time.Since(t)
		fmt.Println("used time: ", used)
	}
}

func afterMware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		fmt.Println("doing after handler^^")
	}
}

func main()  {
	r := gin.Default()
	r.Use(usedTime())
	r.Use(afterMware())
	r.GET("/mw", func(context *gin.Context) {
		// 从gin上下文获取
		user, exists := context.Get("user")
		if exists == true {
			fmt.Println("user: ", user)
		}
		context.JSON(200, gin.H{
			"msg": "success",
		})
	})
	r.Run(":8888")
}
