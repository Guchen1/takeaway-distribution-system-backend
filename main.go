package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type message struct {
	Num   string `json:"num"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Phone string `json:"phone"`
}

func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	go func() {
		var temp message
		time.Sleep(10 * time.Second)
		temp.Num = "1"
		temp.Name = "顾晨"
		temp.Value = "gc"
		temp.Phone = "7196"
		ws.WriteJSON(temp)
		time.Sleep(10 * time.Second)
		temp.Num = "2"
		temp.Name = "吴刘康"
		temp.Value = "wlk"
		temp.Phone = "9303"
		ws.WriteJSON(temp)
	}()
	defer ws.Close()
	for {
		time.Sleep(10 * time.Second)
	}
}

func main() {
	bindAddress := "localhost:2303"
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run(bindAddress)
}
