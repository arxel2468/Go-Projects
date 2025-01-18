package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	tasks      []Task
	mutex      sync.Mutex
	taskFile   = "tasks.json"
	nextTaskID = 1
)

func main() {
	// Load tasks from file
	loadTasks()

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/add", handleAddTask)
	http.HandleFunc("/delete", handleDeleteTask)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		http.Error(w, "Unable to load templates", http.StatusInternalServerError)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	tmpl.Execute(w, tasks)
}

func handleAddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	name := r.FormValue("task")
	if name == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	task := Task{ID: nextTaskID, Name: name}
	tasks = append(tasks, task)
	nextTaskID++
	saveTasks()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i, task := range tasks {
		if task.ID == atoi(id) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	saveTasks()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func atoi(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}

func saveTasks() {
	file, _ := json.MarshalIndent(tasks, "", "  ")
	ioutil.WriteFile(taskFile, file, 0644)
}

func loadTasks() {
	file, err := os.Open(taskFile)
	if err != nil {
		return
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &tasks)

	if len(tasks) > 0 {
		nextTaskID = tasks[len(tasks)-1].ID + 1
	}
}
