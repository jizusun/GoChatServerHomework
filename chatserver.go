package main

import (
	"fmt"
	"net/http"
)

func main() {
	store := &Store{
		Utils: Utilities{},
		Users: make(map[string]bool),
	}
	http.HandleFunc("/message", CreateMessageHandler(store))
	http.HandleFunc("/messages", ReadMessageHandler(store))
	http.HandleFunc("/users", GetUsersHandler(store))
	http.HandleFunc("/status", StatusHandler)
	fmt.Println("Listening on :8081")
	err := http.ListenAndServe(":8081", nil)
	panic(err)
}
