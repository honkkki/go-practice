package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func usedTime() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		// 需要再next之前设置
		context.Set("user", "irene")
		context.Next() // 执行完其他处理函数后执行
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

func main() {
	r := gin.Default()
	g := r.Group("/shop", func(context *gin.Context) {
		fmt.Println("func1 before")
		context.Next()
		fmt.Println("func1 after")
	})

	g.Use(func(context *gin.Context) {
		fmt.Println("func2 before")
		context.Next()
		fmt.Println("func2 after")
	})

	{
		g.GET("/index", func(context *gin.Context) {
			fmt.Println("func3")
			context.JSON(200, gin.H{
				"msg": "ok",
			})
		}, func(context *gin.Context) {
			fmt.Println("func4")
		})
	}

	//r.Use(usedTime())
	//r.Use(afterMware())
	//r.GET("/mw", func(context *gin.Context) {
	//	// 从gin上下文获取
	//	user, exists := context.Get("user")
	//	if exists == true {
	//		fmt.Println("user: ", user)
	//	}
	//	context.JSON(200, gin.H{
	//		"msg": "success",
	//	})
	//})

	log.Fatal(r.Run(":8888"))
}
