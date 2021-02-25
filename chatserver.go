package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jizusun/GoChatServer/handlers"
)

// Message the message is processed by the server a unix timestamp is recorded with each message.
type Message struct {
	User      string `json:"user"`
	Text      string `json:"text"`
	Timestamp int64
}

func NewMessage(user string, text string) Message {
	now := time.Now()
	return Message{
		User:      user,
		Text:      text,
		Timestamp: now.Unix(),
	}
}

// ResponsePostMessage the response for `POST` `/message`
type ResponsePostMessage struct {
	Ok bool `json:"ok"`
}

func main() {
	var messages []Message

	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		var m Message
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		messages = append(messages, m)
		res := ResponsePostMessage{Ok: true}
		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")
		enc.Encode(res)
	})

	http.HandleFunc("/status", handlers.StatusHandler)
	err := http.ListenAndServe(":8081", nil)
	panic(err)
}
