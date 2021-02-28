// Author: Jizu Sun (sunjizu@gmail.com)
package main

import (
	"fmt"
	"net/http"
)

func main() {
	store := &messageStore{
		Utils: utilities{},
	}
	http.HandleFunc("/message", createMessageHandler(store))
	http.HandleFunc("/messages", readMessageHandler(store))
	http.HandleFunc("/users", getUsersHandler(store))
	http.HandleFunc("/status", statusHandler)
	fmt.Println("Listening on :8081")

	err := http.ListenAndServe(":8081", nil)
	panic(err)
}
