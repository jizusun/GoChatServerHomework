package main

import (
	"encoding/json"
	"net/http"
)

// ResponsePostMessage the response for `POST` `/message`
type ResponsePostMessage struct {
	Ok bool `json:"ok"`
}

var messages []*Message

// ReadMessage read the last 100 messages, sorted ascending by timestamp
func ReadMessage(w http.ResponseWriter, r *http.Request) {
	results := Store{
		Messages: messages,
	}
	output, err := json.MarshalIndent(&results, "", "  ")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	messages = append(messages, &m)
	res := ResponsePostMessage{Ok: true}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}
