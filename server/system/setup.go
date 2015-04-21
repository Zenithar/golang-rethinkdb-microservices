package system

import (
	r "github.com/dancannon/gorethink"

	"zenithar.org/microservices/server/domain"
	"zenithar.org/microservices/server/env"
)

func Setup() error {

	// Initialize RethinkDB connection
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
	if err != nil {
		return err
	}

	// Put the session in the environment package
	env.Session = session

	// Initialize DAO Repositories
	env.Todos = &domain.TodoRepository{
		Session: session,
	}

	return nil
}
