package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gopkg.in/olahol/melody.v1"
)

var (
	port     = ":7788"
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func ws(ctx *gin.Context) {
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()
	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		log.Println(string(msg))
		err = ws.WriteMessage(msgType, msg)
		if err != nil {
			break
		}
	}
}

func main() {
	e := gin.Default()
	m := melody.New()

	e.GET("/", func(ctx *gin.Context) {
		http.ServeFile(ctx.Writer, ctx.Request, "index.html")
	})

	e.GET("/ws", func(ctx *gin.Context) {
		m.HandleRequest(ctx.Writer, ctx.Request)
	})

	m.HandleMessage(func(session *melody.Session, msg []byte) {
		err := m.Broadcast(msg)
		if err != nil {
			return
		}
	})
	//e.GET("/ws", ws)

	log.Fatal(e.Run(port))
}
