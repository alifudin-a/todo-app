package api

import (
	"net/http"

	db "github.com/alifudin-a/todo-app/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createTaskReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Complete    string `json:"complete" binding:"required"`
}

// CreateTask create a new task
func (server *Server) createTask(ctx *gin.Context) {
	var req createTaskReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTaskParams{
		Title:       req.Title,
		Description: req.Description,
		Complete:    req.Complete,
	}

	task, err := server.todo.CreateTask(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"task": task})
}
