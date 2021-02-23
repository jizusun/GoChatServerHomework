package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/status",
		func(c http.ResponseWriter, req *http.Request) {
			c.Write([]byte("alive"))
		})
	err := http.ListenAndServe(":8081", nil)
	panic(err)
}
