package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	v *viper.Viper
)

type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func UserRegister(c *gin.Context) {
	user := new(User)
	err := c.ShouldBindJSON(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	log.Println(user.Username)
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}

func main() {
	v = viper.New()
	v.SetConfigFile("./config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("config file not found:", err)
		}
		log.Fatal(err)
	}
	v.WatchConfig()

	port := v.GetString("http.port")

	e := gin.Default()
	e.GET("/config", func(context *gin.Context) {
		port := v.GetString("http.port")
		name := v.GetString("http.name")
		v.Set("http.ip", "127.0.0.1")
		err := v.WriteConfig()
		if err != nil {
			log.Println(err)
		}
		context.String(http.StatusOK, "%v running on :%v", name, port)
		return
	})

	e.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * 5)
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	e.POST("/register", UserRegister)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%v", port),
		Handler:      e,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Println("http server error:", err)
		}
	}()

	// graceful shut down
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("server shutting down...")
	err = s.Shutdown(context.Background())
	if err != nil {
		log.Println("shut down error:", err)
		return
	}

	log.Println("server exit")
}
