package main

import (
	"fmt"
	"os"

	handlers "github.com/alifudin-a/todo-app/handlers/todo"
	"github.com/alifudin-a/todo-app/util"
	"github.com/gin-gonic/gin"
)

func main() {
	util.LoadEnv()
	r := gin.Default()

	r.GET("/todo", handlers.GetTodoList)
	r.POST("/todo", handlers.AddTodo)
	r.DELETE("/todo/:id", handlers.DeleteTodo)
	r.PUT("/todo", handlers.CompleteTodo)

	err := r.Run(":" + os.Getenv("TODO_APP_ADDRESS"))
	if err != nil {
		fmt.Println("Unable to start ", err)
	}
}
