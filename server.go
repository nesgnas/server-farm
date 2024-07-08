package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	// Solve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func ws(c *gin.Context, msgCh chan string) {
	// Upgrade get request to WebSocket protocol
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()

	go func() {
		for {
			// Read data from ws
			mt, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)

			// Process client message
			switch string(message) {
			case "Ping":
				msgCh <- "pong"
			default:
				msgCh <- "Unknown command"
			}

			// Echo message back to client
			err = ws.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}()

	for {
		select {
		case msg := <-msgCh:
			err := ws.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}

func main() {
	fmt.Println("Websocket Server!")

	
	r := gin.Default()

	// Create a channel for console messages
	msgCh := make(chan string)

	//// Start a goroutine to read input from console
	//go func() {
	//	reader := bufio.NewReader(os.Stdin)
	//	for {
	//		fmt.Print("Enter message: ")
	//		text, err := reader.ReadString('\n')
	//		if err != nil {
	//			log.Println("Error reading from console:", err)
	//			continue
	//		}
	//		text = strings.TrimSpace(text)
	//		msgCh <- text
	//	}
	//}()

	r.GET("/ws", func(c *gin.Context) {
		ws(c, msgCh)
	})

	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{"ping":"pong"})
	})
	r.Run(":8000")
}
