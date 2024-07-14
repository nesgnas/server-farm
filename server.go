package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server-farm/controler"
	"server-farm/router"
)

func main() {
	fmt.Println("Websocket Server!")

	r := gin.Default()

	router.WebSocket(r)

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
		controler.Ws(c, msgCh, "")
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})
	r.Run(":8000")
}
