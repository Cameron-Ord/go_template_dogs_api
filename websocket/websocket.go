package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"main.go/requests"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		message := string(p)
		fmt.Printf("Received message: %s, %d\n", message, messageType)
		if message == "Give me DOGS" {
			dog_images, err := requests.Req_Dogs()
			if err != nil {
				return
			}

			json_images, err := json.Marshal(dog_images)
			if err != nil {
				return
			}
			err = conn.WriteMessage(1, json_images)
			if err != nil {
				return
			}

		}
	}
}
