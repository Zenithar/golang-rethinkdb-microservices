.PHONY: deps

default: build

deps:
	go get github.com/tools/godep
	go get github.com/gogo/protobuf/proto
	go get github.com/gogo/protobuf/protoc-gen-gogo
	go get github.com/gogo/protobuf/gogoproto
	go get google.golang.org/grpc

build:
	go build -o srv_server ./server
	go build -o srv_client ./client
	go build -o srv_rest ./rest
