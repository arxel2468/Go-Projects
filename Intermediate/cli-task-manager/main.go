package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("CLI Task Manager")
	fmt.Println("Usage:")
	fmt.Println("  add <task description> - Add a new task")
	fmt.Println("  list                   - List all tasks")
	fmt.Println("  done <task ID>         - Mark task as done")
	fmt.Println("  delete <task ID>       - Delete a task")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task description required.")
			return
		}
		description := os.Args[2]
		err := AddTask(description)
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		fmt.Println("Task added successfully.")
	case "list":
		err := ListTasks()
		if err != nil {
			fmt.Println("Error listing tasks:", err)
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task ID required.")
			return
		}
		taskID := os.Args[2]
		err := MarkTaskDone(taskID)
		if err != nil {
			fmt.Println("Error marking task as done:", err)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Task ID required.")
			return
		}
		taskID := os.Args[2]
		err := DeleteTask(taskID)
		if err != nil {
			fmt.Println("Error deleting task:", err)
		}
	default:
		printUsage()
	}
}
