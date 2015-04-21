package main

import (
	"code.google.com/p/go.net/context"

	"github.com/asim/go-micro/cmd"
	"github.com/asim/go-micro/server"
	log "github.com/golang/glog"

	"zenithar.org/microservices/server/env"
	todo "zenithar.org/microservices/server/proto/todo"
	"zenithar.org/microservices/server/system"
)

type Todo struct{}

func (self *Todo) Create(ctx context.Context, r *todo.CreateTodoRequest, w *todo.CreateTodoResponse) error {
	t, err := env.Todos.Create(r)

	w.Todo = &todo.Todo{
		Id:   &t.ID,
		Task: t.Task,
		Done: t.Done,
		Date: t.Date.Unix(),
	}

	return err
}

func (self *Todo) List(ctx context.Context, r *todo.GetTodoListRequest, w *todo.GetTodoListResponse) error {
	todos, err := env.Todos.List(r)

	for _, t := range todos {
		w.Members = append(w.Members, &todo.Todo{
			Id:   &t.ID,
			Task: t.Task,
			Done: t.Done,
			Date: t.Date.Unix(),
		})
	}

	return err
}

func main() {
	// optionally setup command line usage
	cmd.Init()

	// Initialize database connection
	err := system.Setup()
	if err != nil {
		log.Fatalln(err.Error())
	}

	server.Name = "go.micro.srv.todo"

	// Initialise Server
	server.Init()

	// Register Handlers
	server.Register(
		server.NewReceiver(
			new(Todo),
		),
	)

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
