// Author: Jizu Sun
package main

// Message the message is processed by the server a unix timestamp is recorded with each message.
type Message struct {
	Timestamp int64    `json:"timestamp"`
	User      username `json:"user"`
	Text      string   `json:"text"`
}

type username = string

type messageList struct {
	Messages []*Message `json:"messages"`
}

type messageStore struct {
	messageList
	Utils utilitiesInterface
}

func (s *messageStore) addMessage(m *Message) *Message {
	m.Timestamp = s.Utils.GetTimestamp()
	s.Messages = append(s.Messages, m)
	return m
}

func (s *messageStore) getMessages() []*Message {
	size := len(s.Messages)
	return s.Messages[size-100:]
}

func (s *messageStore) getUsers() []username {
	users := make([]username, 0)
	usersmap := make(map[username]bool)
	for _, m := range s.Messages {
		name := m.User
		_, ok := usersmap[name]
		if !ok {
			usersmap[name] = true
			users = append(users, name)
		}
	}
	return users
}
