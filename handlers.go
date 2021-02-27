package main

import (
	"encoding/json"
	"net/http"
)

// CreationSuccess the response for `POST` `/message`
type CreationSuccess struct {
	Ok bool `json:"ok"`
}

// ReadMessage read the last 100 messages, sorted ascending by timestamp
func ReadMessage(w http.ResponseWriter, r *http.Request, s *Store) {
	results := MessageList{
		Messages: s.Messages,
	}
	output, err := json.MarshalIndent(&results, "", "  ")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func ReadMessageHandler(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ReadMessage(w, r, s)
	}
}

func CreateMessage(w http.ResponseWriter, r *http.Request, s *Store) {
	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	s.AddMessage(&m)
	res := CreationSuccess{Ok: true}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	enc.Encode(res)
}

func CreateMessageHandler(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CreateMessage(w, r, s)
	}
}

// StatusHandler test
func StatusHandler(c http.ResponseWriter, req *http.Request) {
	c.Write([]byte("alive"))
}
