build:
	go build -o chatserver

test:
	go test

run:
	./chatserver

alive:
	curl "http://localhost:8081/status"
	curl -X POST -d "{\"user\": \"batman\", \"text\":\"hello\"}" "http://localhost:8081/message"
	curl "http://localhost:8081/messages"
	curl "http://localhost:8081/users"

