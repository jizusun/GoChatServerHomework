package main

// Message the message is processed by the server a unix timestamp is recorded with each message.
type Message struct {
	User      Username `json:"user"`
	Text      string   `json:"text"`
	Timestamp int64    `json:"timestamp"`
}

type Username = string

type MessageList struct {
	Messages []*Message `json:"messages"`
}

// Store the response of reading messages
type Store struct {
	MessageList
	Utils ExternalInterface
	Users map[Username]bool
}

func (s *Store) AddMessage(m *Message) *Message {
	m.Timestamp = s.Utils.GetTimestamp()
	s.Messages = append(s.Messages, m)
	s.Users[m.User] = true
	return m
}

func (s *Store) GetMessages() []*Message {
	size := len(s.Messages)
	return s.Messages[size-100:]
}

func (s *Store) GetUsers() []Username {
	users := make([]Username, 0, len(s.Users))
	for u := range s.Users {
		users = append(users, u)
	}
	return users
}
