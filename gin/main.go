package main

import (
	"github.com/gin-gonic/gin"
)

func handler(c *gin.Context)  {
	c.JSON(200, gin.H{
		"code": 10000,
		"message": "hello gin",
	})
	
}

func main()  {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", handler)
	r.Run(":9999")
}
