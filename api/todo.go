package api

import (
	"database/sql"
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

type listTasksReq struct {
	Limit  int32 `form:"limit" binding:"required,min=1"`
	Offset int32 `form:"offset" binding:"required,min=1,max=10"`
}

func (server *Server) listTasks(ctx *gin.Context) {
	var req listTasksReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTasksParams{
		Limit:  req.Limit,
		Offset: (req.Offset - 1) * req.Limit,
	}

	list, err := server.todo.ListTasks(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"list_tasks": list})
}

type getTaskReq struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTask(ctx *gin.Context) {
	var req getTaskReq
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	task, err := server.todo.GetTask(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"task": task})
}

type deleteTaskReq struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteTask(ctx *gin.Context) {
	var req deleteTaskReq
	err := ctx.ShouldBindUri(&req)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.todo.DeleteTask(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"task": "Deleted!"})
}

type updateTaskReq struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    string `json:"complete"`
}

func (server *Server) updateTask(ctx *gin.Context) {
	var req updateTaskReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateTaskParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Complete:    req.Complete,
	}

	task, err := server.todo.UpdateTask(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"task": task})
}
