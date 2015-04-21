//go:generate protoc -I. -I$GOPATH/src -I$GOPATH/src/github.com/gogo/protobuf/protobuf --gogo_out=. todo/todo.proto

package proto
