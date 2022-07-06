package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func postFile(c *gin.Context) {
	file, err := c.FormFile("file")
	//form := c.MultipartForm()
	//files := form.File["file"]

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 10005,
			"data": err.Error(),
		})
	}
	dst := fmt.Sprintf("./src/%s", file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 10005,
			"data": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code": 10000,
		"data": "upload success! " + file.Filename,
	})

}

func handler(c *gin.Context) {
	name := c.DefaultQuery("name", "gin")
	c.JSON(200, gin.H{
		"code":    10000,
		"message": "hello " + name,
	})
}

// map转http json
func retJson(c *gin.Context) {
	m := make(map[string]interface{})
	m["name"] = "irene"
	c.JSON(200, m)
}

// 结构体转http json
func struct2Json(c *gin.Context) {
	var user User
	user.Name = "irene"
	user.Age = 20
	c.JSON(http.StatusOK, user)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Static("/img", "./src")
	r.GET("/hello", handler)
	r.POST("/fileUpload", postFile)
	r.GET("/json", retJson)
	r.GET("/s2json", struct2Json)
	r.Run(":9999")
}
