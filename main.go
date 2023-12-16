package main

import (
	"fmt"
	"net/http"

	"main.go/rendering"
	"main.go/websocket"
)

func main() {
	port := "8080"
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	http.HandleFunc("/ws", websocket.HandleWebSocket)
	http.HandleFunc("/", rendering.DefaultPage)
	fmt.Printf("server running on :%s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
