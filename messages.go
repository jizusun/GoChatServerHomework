package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Message the message is processed by the server a unix timestamp is recorded with each message.
type Message struct {
	User      string `json:"user"`
	Text      string `json:"text"`
	Timestamp int64
}

// Messages the response of reading messages
type Messages struct {
	Messages []Message `json:"messages"`
}

func updateTimestamp(m *Message) *Message {
	now := time.Now()
	m.Timestamp = now.Unix()
	return m
}

// ResponsePostMessage the response for `POST` `/message`
type ResponsePostMessage struct {
	Ok bool `json:"ok"`
}

var messages []Message

// ReadMessage read the last 100 messages, sorted ascending by timestamp
func ReadMessage(w http.ResponseWriter, r *http.Request) {
	results := Messages{
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
	updateTimestamp(&m)
	messages = append(messages, m)
	res := ResponsePostMessage{Ok: true}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}
