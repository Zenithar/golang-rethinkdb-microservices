package main

import (
	"fmt"
	"time"

	"github.com/asim/go-micro/client"
	todo "zenithar.org/microservices/server/proto/todo"
)

func main() {
	// Create new request to service go.micro.service.go-template, method Example.Call
	req := client.NewRequest("go.micro.srv.todo", "Todo.Create", &todo.CreateTodoRequest{
		Todo: &todo.Todo{
			Task: "Do a barrel roll !",
			Done: false,
			Date: time.Now().Unix(),
		},
	})
	rsp := &todo.CreateTodoResponse{}

	// Call service
	if err := client.Call(req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%V\n", rsp.GetTodo())
}
