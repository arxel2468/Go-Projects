package main

import (
	"distributed-task-scheduler/api"
	"distributed-task-scheduler/scheduler"
	"log"
)

func main() {
	// Initialize the scheduler
	taskScheduler := scheduler.NewScheduler()

	// Initialize the API server
	router := api.SetupRouter(taskScheduler)

	// Start the server
	log.Println("Server running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
