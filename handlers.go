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
	w.Header().Set("Content-Type", "application/json")
	output, _ := json.MarshalIndent(&results, "", "  ")
	_, err := w.Write(output)
	if err != nil {
		return
	}
}

func ReadMessageHandler(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ReadMessage(w, r, s)
	}
}

type UserList struct {
	Users []Username `json:"users"`
}

func GetUsers(w http.ResponseWriter, r *http.Request, s *Store) {
	users := s.GetUsers()
	res := UserList{
		Users: users,
	}
	w.Header().Set("Content-Type", "application/json")
	output, _ := json.MarshalIndent(&res, "", "  ")
	_, err := w.Write(output)
	if err != nil {
		return
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
	w.Header().Set("Content-Type", "application/json")
	output, _ := json.MarshalIndent(&res, "", "  ")
	_, err = w.Write(output)
	if err != nil {
		return
	}
	// enc := json.NewEncoder(w)
	// enc.SetIndent("", "  ")
	// enc.Encode(res)
}

func CreateMessageHandler(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CreateMessage(w, r, s)
	}
}

func GetUsersHandler(s *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		GetUsers(w, r, s)
	}
}

// StatusHandler test
func StatusHandler(c http.ResponseWriter, req *http.Request) {
	output := []byte("alive")
	_, err := c.Write(output)
	if err != nil {
		return
	}
}
