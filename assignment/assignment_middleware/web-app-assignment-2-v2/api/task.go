package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskRepo service.TaskService) *taskAPI {
	return &taskAPI{taskRepo}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.taskService.Store(&newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add task success"})
}

func (t *taskAPI) UpdateTask(c *gin.Context) {
	// TODO: answer here
	_, err := c.Cookie("session_token")
	if err != nil {
		// Jika tidak ada cookie session_token, kirim respon HTTP 401 Unauthorized
		c.AbortWithStatus(http.StatusUnauthorized)
		return

	}

	taskIDStr := c.Param("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	var updatedTask model.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	updatedTask.ID = taskID
	err = t.taskService.Update(&updatedTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "update task success"})
}

func (t *taskAPI) DeleteTask(c *gin.Context) {
	// TODO: answer here
	_, err := c.Cookie("session_token")
	if err != nil {
		// Jika tidak ada cookie session_token, kirim respon HTTP 401 Unauthorized
		c.AbortWithStatus(http.StatusUnauthorized)
		return

	}
	taskIDStr := c.Param("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	err = t.taskService.Delete(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "delete task success"})
}

func (t *taskAPI) GetTaskByID(c *gin.Context) {
	_, err := c.Cookie("session_token")
	if err != nil {
		// Jika tidak ada cookie session_token, kirim respon HTTP 401 Unauthorized
		c.AbortWithStatus(http.StatusUnauthorized)
		return

	}
	taskIDStr := c.Param("id")
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := t.taskService.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	if task == nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskAPI) GetTaskList(c *gin.Context) {
	// TODO: answer here
	_, err := c.Cookie("session_token")
	if err != nil {
		// Jika tidak ada cookie session_token, kirim respon HTTP 401 Unauthorized
		c.AbortWithStatus(http.StatusUnauthorized)
		return

	}
	taskList, err := t.taskService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, taskList)
}

func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	// TODO: answer here
	_, err := c.Cookie("session_token")
	if err != nil {
		// Jika tidak ada cookie session_token, kirim respon HTTP 401 Unauthorized
		c.AbortWithStatus(http.StatusUnauthorized)
		return

	}
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid category ID"})
		return
	}

	taskList, err := t.taskService.GetTaskCategory(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, taskList)
}
