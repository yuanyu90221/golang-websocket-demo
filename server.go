package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	httpServer := gin.Default()
	// upgrader := websocket.Upgrader{}

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
	httpServer.GET("/echo", func(ctx *gin.Context) {
		// do switch protocol
		websocketConn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			log.Fatal(err)
			return
		}
		for {
			// Read message from browser
			msgType, msg, err := websocketConn.ReadMessage()
			if err != nil {
				log.Fatal(err)
				return
			}
			log.Printf("%s sent: %s\n", websocketConn.RemoteAddr(), string(msg))
			// Write message backe to browser
			if err = websocketConn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})
	httpServer.StaticFile("/", "websockets.html")
	err := http.ListenAndServe(":8080", httpServer)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("websocket server listen on 8080")
}
