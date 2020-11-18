package api

import (
	db "github.com/alifudin-a/todo-app/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests
type Server struct {
	todo   *db.Todo
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(todo *db.Todo) *Server {
	server := &Server{todo: todo}
	router := gin.Default()

	router.POST("/todo", server.createTask)
	router.GET("/todo", server.listTasks)

	server.router = router
	return server
}

// Start runs HTTP server on a spesific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// ErrorResponse response for error
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
