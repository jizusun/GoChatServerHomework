# GoChatServer

[![build](https://github.com/jizusun/GoChatServerHomework/actions/workflows/build.yml/badge.svg)](https://github.com/jizusun/GoChatServerHomework/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/jizusun/GoChatServerHomework/branch/main/graph/badge.svg?token=UB5E9NIZ68)](https://codecov.io/gh/jizusun/GoChatServerHomework)

The goal of the exercise is to implement a chat server. 
The chat server will run on http://localhost:8081 and will support the following REST API: 

## 1. `GET /messages`

list 100 most recent messages, sorted by 'timestamp' posted to the chat server.

example:
=======

```
curl -H "Content-Type: application/json" http://localhost:8081/messages

{
  "messages: [
    {"timestamp": 1491345710.18, "user": "superman", "text": "hello"},
    {"timestamp": 1491345713.18, "user": "batman", "text": "hello"}
  ]
}

```

## 2. `POST /message`

a request to post the given message. 
when the message is processed by the server a unix timestamp is recorded with each message.

example:
========

```
curl -X POST -H "Content-Type: application/json" --data '{"user":"superman", "text":"hello"}' http://localhost:8081/message

{
  "ok": true
}
```

## 3. `GET /users`

a request to return a set of users that have posted messages so far.

example:
========

```
curl -H "Content-Type: application/json" http://localhost:8081/users

{
  "users": [
    "superman", "batman"
  ]
}
```

the server should respond with 404 to all other requests not listed above
 
Instructions:
=============

1. Preferred programming language: Go
2. Please provide a *single implementation* with no external dependencies.
3. The submitted implementation will be tested by an automated script 
      that will build the chat app with: `make build` 
      and execute it with: `make run`
    (see attached Makefile)
4. Please describe what metrics you would monitor to track the performance of the chat server.
5. Bonus Points: please describe how you would improve the chat server API.

