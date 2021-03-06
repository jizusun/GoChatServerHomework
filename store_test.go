package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeUtilities struct{}

func (ex FakeUtilities) GetTimestamp() int64 {
	return 149134571
}

func TestStore_AddMessage_GetUsers(t *testing.T) {
	s := messageStore{
		Utils: FakeUtilities{},
	}
	m := &Message{
		User: "superman",
		Text: "hello",
	}
	expected := &Message{
		User:      "superman",
		Text:      "hello",
		Timestamp: 149134571,
	}
	actual := s.addMessage(m)
	assert.Equal(t, expected, actual)
	assert.Equal(t, s.Messages[0], expected)
	assert.Equal(t, len(s.Messages), 1)

	s.addMessage(&Message{
		User: "batman",
		Text: "hello",
	})
	s.addMessage(&Message{
		User: "batman",
		Text: "hello world",
	})
	actualUsers := s.getUsers()

	assert.Equal(t, actualUsers, []username{"superman", "batman"})
}

func TestStore_GetMessages(t *testing.T) {
	var messages []*Message
	size := 104
	for i := 1; i < size+1; i++ {
		id := strconv.Itoa(i)
		m := &Message{
			User:      "user" + id,
			Text:      "text" + id,
			Timestamp: int64(i),
		}
		messages = append(messages, m)
	}
	s := messageStore{
		Utils: FakeUtilities{},
		messageList: messageList{
			Messages: messages,
		},
	}
	actual := s.getMessages()
	lastMessage := messages[size-1]
	assert.Equal(t, len(messages), size)
	assert.Equal(t, len(actual), 100)
	assert.Equal(t, actual[len(actual)-1], lastMessage)
	assert.Equal(t, actual[0], messages[size-100])
}
