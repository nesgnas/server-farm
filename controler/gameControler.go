package controler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	// Solve cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

func Ws(c *gin.Context, msgCh chan string, method string) {
	// Upgrade get request to WebSocket protocol
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer ws.Close()

	switch method {
	case "":
		go func() {
			for {
				// Read data from ws
				_, message, err := ws.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					break
				}
				log.Printf("recv: %s", message)

				// Process client message
				switch string(message) {
				case "Ping":
					msgCh <- "pong"
					break
				default:
					msgCh <- string(message)
				}

				//// Echo message back to client
				//err = ws.WriteMessage(mt, message)
				//if err != nil {
				//	log.Println("write:", err)
				//	break
				//}
			}
		}()
		break
	case "InventoryUpdate":
		go func() {
			for {
				// Read data from ws
				_, message, err := ws.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					break
				}
				log.Printf("recv: %s", message)

				// Process client message
				switch string(message) {
				case "Ping":
					msgCh <- "pong"
					break
				case "GetInventory":
					go func() {
						if err := SentJson(ws); err != nil {
							log.Println("SentJson error:", err)
						}
					}()
				default:
					msgCh <- string(message)
				}

				//// Echo message back to client
				//err = ws.WriteMessage(mt, message)
				//if err != nil {
				//	log.Println("write:", err)
				//break
				//}
			}
		}()
		break
	}

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

// function sent inventory12.json in jsonData directory through websocket
func SentJson(ws *websocket.Conn) error {
	filePath := "jsonData/inventory12.json"

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = ws.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		return err
	}

	return nil
}

func InventoryUpdate() gin.HandlerFunc {
	msgCh := make(chan string)
	return func(c *gin.Context) {
		Ws(c, msgCh, "InventoryUpdate")
	}
}
