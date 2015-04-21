package env

import (
	r "github.com/dancannon/gorethink"
	"zenithar.org/microservices/server/domain"
)

var (
	// Session stores the RethinkDB connection
	Session *r.Session

	// Todos is the database access layer to manage todos table
	Todos *domain.TodoRepository
)
