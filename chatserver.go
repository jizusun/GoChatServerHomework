package main

import (
	"net/http"
)

func main() {
	// http.HandleFunc("/message", CreateMessage)
	// http.HandleFunc("/messages", ReadMessage)
	http.HandleFunc("/status", StatusHandler)
	err := http.ListenAndServe(":8081", nil)
	panic(err)
}
