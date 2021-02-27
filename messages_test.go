package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreyvit/diff"
)

// MessageReadHandler
// should list 100 most recent messages sorted by 'timestamp'

// `POST message` MessageCreateHandler
// should return indented {"ok": true}
// should record a unix timestamp with each message

func createMessage() {

}

func TestMessageCreateHandler(t *testing.T) {
	jsonStr := []byte(`{"user":"superman", "text":"hello"}`)
	req, err := http.NewRequest("POST", "/message", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateMessage)
	handler.ServeHTTP(rr, req)
	expected := `{
  "ok": true
}
`
	actual := rr.Body.String()
	if actual != expected {
		t.Errorf("Result not as expected:\n%v", diff.LineDiff(expected, actual))
		// t.Errorf("handler returned unexpected body: got %v want %v",
		// 	rr.Body.String(), expected)
	}
}
