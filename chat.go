package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//聊天页面
func chat(c *gin.Context) {

	c.HTML(200, "chat.html", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//接收消息 发送消息

func OnWebSocket(c *gin.Context) {
	var err error
	var conn *websocket.Conn
	conn, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	ConnMap[c.Param("uname")] = conn

	var t int
	var msg []byte
	for {
		t, msg, err = conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("从客户端收到:", string(msg))

		broadcase(t, msg)
	}

	defer conn.Close()

}

//广播
func broadcase(t int, msg []byte) {
	for _, v := range ConnMap {

		v.WriteMessage(t, msg)
	}
}
