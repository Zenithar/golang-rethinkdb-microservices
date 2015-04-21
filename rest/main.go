package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asim/go-micro/client"
	"github.com/zenazn/goji"
	todo "zenithar.org/microservices/server/proto/todo"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	// Create new request to service go.micro.service.go-template, method Example.Call
	req := client.NewRequest("go.micro.srv.todo", "Todo.List", &todo.GetTodoListRequest{})
	rsp := &todo.GetTodoListResponse{}

	// Call service
	if err := client.Call(req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	payload, _ := json.Marshal(rsp)

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func main() {
	goji.Get("/todos", GetTodos)
	goji.Serve()
}
