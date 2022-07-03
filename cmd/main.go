package main

import (
	"fmt"
	"net/http"

	controllers "github.com/fuadnafiz98/go-postgres/internal/api/controllers"
	database "github.com/fuadnafiz98/go-postgres/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	server := gin.Default()

	database.ConnectDatabase()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})

	server.GET("/ws", func(c *gin.Context) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer ws.Close()

		for {
			// read message from client.
			mt, message, err := ws.ReadMessage()
			println(string(message))

			if err != nil {
				fmt.Println(err)
				return
			}

			if string(message) == "change" {
				message = []byte("changed")
			}

			// send message to the client.
			err = ws.WriteMessage(mt, message)

			if err != nil {
				fmt.Println(err)
				return
			}
		}
	})

	server.GET("/messages", controllers.GetAllMessages)

	server.Run()
}
