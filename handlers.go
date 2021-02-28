// Author: Jizu Sun (sunjizu@gmail.com)
package main

import (
	"encoding/json"
	"net/http"
)

type creationSuccess struct {
	Ok bool `json:"ok"`
}

func readMessage(w http.ResponseWriter, _ *http.Request, s *messageStore) {
	results := messageList{
		Messages: s.Messages,
	}

	w.Header().Set("Content-Type", "application/json")

	output, _ := json.MarshalIndent(&results, "", "  ")
	if _, err := w.Write(output); err != nil {
		return
	}
}

func readMessageHandler(s *messageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		readMessage(w, r, s)
	}
}

type userList struct {
	Users []username `json:"users"`
}

func getUsers(w http.ResponseWriter, _ *http.Request, s *messageStore) {
	users := s.getUsers()
	res := userList{
		Users: users,
	}

	w.Header().Set("Content-Type", "application/json")

	output, _ := json.MarshalIndent(&res, "", "  ")
	if _, err := w.Write(output); err != nil {
		return
	}
}

func createMessage(w http.ResponseWriter, r *http.Request, s *messageStore) {
	var m Message
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	s.addMessage(&m)

	res := creationSuccess{Ok: true}

	w.Header().Set("Content-Type", "application/json")

	output, _ := json.MarshalIndent(&res, "", "  ")
	if _, err := w.Write(output); err != nil {
		return
	}
}

func createMessageHandler(s *messageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createMessage(w, r, s)
	}
}

func getUsersHandler(s *messageStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getUsers(w, r, s)
	}
}

func statusHandler(c http.ResponseWriter, req *http.Request) {
	output := []byte("alive")
	_, err := c.Write(output)
	if err != nil {
		return
	}
}
