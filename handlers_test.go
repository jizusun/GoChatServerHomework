package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MessageReadHandler
// should list 100 most recent messages sorted by 'timestamp'

// `POST message` MessageCreateHandler
// should return indented {"ok": true}
// should record a unix timestamp with each message

func TestMessageCreateHandler(t *testing.T) {
	store := &Store{
		Utils: External{},
		Users: make(map[string]bool),
	}
	jsonStr := []byte(`{"user":"superman", "text":"hello"}`)
	req, _ := http.NewRequest("POST", "/message", bytes.NewBuffer(jsonStr))
	rec := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.HandleFunc("/message", CreateMessageHandler(store))
	mux.HandleFunc("/messages", ReadMessageHandler(store))
	mux.ServeHTTP(rec, req)

	expectedCreateRes := `{
  "ok": true
}
`
	actualCreateRes := rec.Body.String()
	assert.Equal(t, expectedCreateRes, actualCreateRes)
}

func TestMessageReadHandler(t *testing.T) {
	message1 := Message{
		User:      "superman",
		Text:      "hello",
		Timestamp: int64(1491345710),
	}
	message2 := Message{
		User:      "batman",
		Text:      "hello",
		Timestamp: int64(1491345713),
	}
	store := &Store{
		Utils: External{},
		Users: make(map[string]bool),
		MessageList: MessageList{
			Messages: []*Message{&message1, &message2},
		},
	}
	req, _ := http.NewRequest("GET", "/messages", nil)
	rec := httptest.NewRecorder()
	mux := http.NewServeMux()
	mux.HandleFunc("/messages", ReadMessageHandler(store))
	mux.ServeHTTP(rec, req)

	actualReadRes := rec.Body.String()
	expectedReadRes := `
{
  "messages: [
    {"timestamp": 1491345710, "user": "superman", "text": "hello"},
    {"timestamp": 1491345713, "user": "batman", "text": "hello"}
  ]
}
`
	assert.Equal(t, expectedReadRes, actualReadRes)
}

func TestStatusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusHandler)
	handler.ServeHTTP(rr, req)
	expected := "alive"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
