.PHONY: deps generate build

export PATH := $(GOPATH)/bin:$(PATH)

default: build

deps:
	go get github.com/tools/godep
	go get github.com/gogo/protobuf/proto
	go get github.com/gogo/protobuf/protoc-gen-gogo
	go get github.com/gogo/protobuf/gogoproto
	go get google.golang.org/grpc

generate:
	go generate ./...

build: generate
	go build -o srv_server ./server
	go build -o srv_client ./client
	go build -o srv_rest ./rest
