package main

import (
	"log"
	"os"

	"github.com/alifudin-a/todo-app/api"
	db "github.com/alifudin-a/todo-app/db/sqlc"
	"github.com/alifudin-a/todo-app/util"
	_ "github.com/lib/pq"
)

func main() {
	util.LoadEnv()
	conn := db.OpenDBConn()

	todo := db.NewTodo(conn)
	server := api.NewServer(todo)

	err := server.Start(os.Getenv("TODO_APP_ADDRESS"))
	if err != nil {
		log.Fatal("Unable to start server: ", err)
	}
}
