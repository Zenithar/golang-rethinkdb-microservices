package domain

import (
	"time"

	r "github.com/dancannon/gorethink"
	todo "zenithar.org/microservices/server/proto/todo"
)

type Todo struct {
	ID   string    `gorethink:"id,omitempty"`
	Task string    `gorethink:"task"`
	Done bool      `gorethink:"done"`
	Date time.Time `gorethink:"date"`
}

type TodoRepository struct {
	Session *r.Session
}

func (repo *TodoRepository) List(req *todo.GetTodoListRequest) ([]*Todo, error) {
	var todos []*Todo

	rows, err := r.Db("todo").Table("todos").Run(repo.Session)
	if err != nil {
		return todos, nil
	}

	err = rows.All(&todos)
	if err != nil {
		return todos, err
	}

	return todos, err
}

func (repo *TodoRepository) Create(req *todo.CreateTodoRequest) (*Todo, error) {
	var todo = &Todo{
		Task: req.Todo.GetTask(),
		Done: req.Todo.GetDone(),
		Date: time.Unix(req.Todo.GetDate(), 0),
	}

	res, err := r.Db("todo").Table("todos").Insert(todo).RunWrite(repo.Session)
	if err != nil {
		return nil, err
	}

	if todo.ID == "" && len(res.GeneratedKeys) == 1 {
		todo.ID = res.GeneratedKeys[0]
	}

	return todo, nil
}
