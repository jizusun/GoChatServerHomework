package main

// Message the message is processed by the server a unix timestamp is recorded with each message.
type Message struct {
	Timestamp int64    `json:"timestamp"`
	User      Username `json:"user"`
	Text      string   `json:"text"`
}

type Username = string

type MessageList struct {
	Messages []*Message `json:"messages"`
}

// Store the response of reading messages
type Store struct {
	MessageList
	Utils UtilitiesInterface
}

func (s *Store) AddMessage(m *Message) *Message {
	m.Timestamp = s.Utils.GetTimestamp()
	s.Messages = append(s.Messages, m)
	return m
}

func (s *Store) GetMessages() []*Message {
	size := len(s.Messages)
	return s.Messages[size-100:]
}

func (s *Store) GetUsers() []Username {
	users := make([]Username, 0)
	usersmap := make(map[Username]bool)
	for _, m := range s.Messages {
		username := m.User
		_, ok := usersmap[username]
		if !ok {
			usersmap[username] = true
			users = append(users, username)
		}
	}
	return users
}
