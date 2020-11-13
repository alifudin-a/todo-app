package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/alifudin-a/todo-app/util"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func createTaskTodo(t *testing.T) Task {
	arg := CreateTaskParams{
		Title:       util.RandomTitle(),
		Description: util.RandomDescription(),
		Complete:    "false",
	}

	ctx := context.Background()
	task, err := testQueries.CreateTask(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Description, task.Description)
	require.Equal(t, arg.Complete, task.Complete)
	require.NotZero(t, task.ID)

	return task
}

func TestCreateTask(t *testing.T) {
	createTaskTodo(t)
}

func TestGetTask(t *testing.T) {
	ctx := context.Background()
	createTasks := createTaskTodo(t)
	task, err := testQueries.GetTask(ctx, createTasks.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, createTasks.ID, task.ID)
	require.Equal(t, createTasks.Title, task.Title)
	require.Equal(t, createTasks.Description, task.Description)
	require.Equal(t, createTasks.Complete, task.Complete)
}

func TestListTasks(t *testing.T) {
	createTaskTodo(t)
	arg := ListTasksParams{
		Limit:  5,
		Offset: 0,
	}

	ctx := context.Background()
	tasks, err := testQueries.ListTasks(ctx, arg)
	require.NoError(t, err)
	require.Len(t, tasks, 5)

	for _, task := range tasks {
		require.NotEmpty(t, task)
	}
}

func TestDeleteTask(t *testing.T) {
	createTask := createTaskTodo(t)
	ctx := context.Background()
	err := testQueries.DeleteTask(ctx, createTask.ID)
	require.NoError(t, err)

	task, err := testQueries.GetTask(ctx, createTask.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, task)
}

func TestUpdateTask(t *testing.T) {
	createTask := createTaskTodo(t)

	arg := UpdateTaskParams{
		ID:          createTask.ID,
		Title:       "ID 16 Update",
		Description: "ID 16 Update",
		Complete:    "true",
	}

	ctx := context.Background()

	task, err := testQueries.UpdateTask(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, createTask.ID, task.ID)
	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Description, task.Description)
	require.Equal(t, arg.Complete, task.Complete)

}
