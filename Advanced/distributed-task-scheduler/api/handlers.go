package api

import (
	"distributed-task-scheduler/scheduler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(s *scheduler.Scheduler) gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks := s.GetAllTasks()
		c.JSON(http.StatusOK, tasks)
	}
}

func AddTask(s *scheduler.Scheduler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task scheduler.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := s.AddTask(task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Task added successfully"})
	}
}
