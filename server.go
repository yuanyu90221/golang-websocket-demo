package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	httpServer := gin.Default()
	upgrader := websocket.Upgrader{}

	httpServer.GET("/ws", func(ctx *gin.Context) {
		// do switch protocol
		websocketConn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		_, message, err := websocketConn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("Message: %s\n", string(message))
		err = websocketConn.WriteMessage(websocket.TextMessage, []byte("Hello from server!\n"))
		if err != nil {
			log.Fatal(err)
			return
		}
	})
	err := http.ListenAndServe(":8080", httpServer)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("websocket server listen on 8080")
}
