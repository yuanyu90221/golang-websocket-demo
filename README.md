# golang with websocket

go websocket server example use github.com/gorilla/websocket package for websocket

and use gin as http handle server

## main logic

```go
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
```

## client code for test

```go
package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	url := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello from client!\n"))
	if err != nil {
		log.Fatal(err)
		return
	}
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Received: %s\n", string(message))
}
```

## Makefile

```makefile
build-server:
	@go build -o bin/websocket-server server.go

run-server: build-server
	./bin/websocket-server

build-client:
	@go build -o bin/websocket-client client/client.go

run-client: build-client
	./bin/websocket-client
```