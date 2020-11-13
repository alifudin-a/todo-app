package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/alifudin-a/todo-app/util"
)

// Todo contains global variable
type Todo struct {
	*Queries
	db *sql.DB
}

// NewTodo creates a new Todo
func NewTodo(db *sql.DB) *Todo {
	return &Todo{
		db:      db,
		Queries: New(db),
	}
}

// OpenDBConn open db connection
func OpenDBConn() *sql.DB {
	util.LoadEnv()

	conn, err := sql.Open(os.Getenv("DB_DRIVER"), os.Getenv("DB_SOURCE"))
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	log.Println("Database Successfully Connected!")

	return conn
}
