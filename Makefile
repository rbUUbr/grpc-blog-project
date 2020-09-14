# Run server 
.PHONY: server
server:
	go run server/server.go

# Run client 
.PHONY: client
client:
	go run client/client.go

# Run protoc generator
.PHONY: proto
proto:
	protoc -I=. --go_out=plugins=grpc:. pb/post.proto

