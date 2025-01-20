package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const taskFile = "tasks.json"

func loadTasks() ([]Task, error) {
	file, err := os.ReadFile(taskFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(taskFile, data, 0644)
}

func AddTask(description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	task := Task{
		ID:          fmt.Sprintf("%d", len(tasks)+1),
		Description: description,
		Done:        false,
	}
	tasks = append(tasks, task)

	return saveTasks(tasks)
}

func ListTasks() error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("%s. %s [%s]\n", task.ID, task.Description, status)
	}
	return nil
}

func MarkTaskDone(taskID string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Done = true
			return saveTasks(tasks)
		}
	}

	return errors.New("task not found")
}

func DeleteTask(taskID string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return saveTasks(tasks)
		}
	}

	return errors.New("task not found")
}
