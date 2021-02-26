package main

import (
	"net/http"

	"github.com/jizusun/GoChatServer/handlers"
)

func main() {
	// var messages []Message
	http.HandleFunc("/message", handlers.PostMessageHandler)
	http.HandleFunc("/status", handlers.StatusHandler)
	err := http.ListenAndServe(":8081", nil)
	panic(err)
}
