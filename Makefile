build-server:
	@go build -o bin/websocket-server server.go

run-server: build-server
	./bin/websocket-server

build-client:
	@go build -o bin/websocket-client client/client.go

run-client: build-client
	./bin/websocket-client