package api

import (
	"distributed-task-scheduler/scheduler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(s *scheduler.Scheduler) *gin.Engine {
	router := gin.Default()

	// Serve the frontend
	router.Static("/static", "./frontend")
	router.LoadHTMLGlob("frontend/*.html")
	router.StaticFile("/app.js", "./frontend/app.js")


	// API routes
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/tasks", GetTasks(s))
	router.POST("/tasks", AddTask(s))

	return router
}
