# 📝 To-Do List Application

A simple web-based To-Do List application built with Go to manage daily tasks. It allows users to add and delete tasks with persistence using a JSON file.

## 🚀 Features
- Add new tasks  
- View all tasks  
- Delete tasks  
- Persistent storage using `tasks.json`

## 🛠️ Prerequisites
- [Go](https://golang.org/dl/) installed (version 1.16+ recommended)

## ⚙️ Setup & Run

1. **Clone the Repository**  
   ```bash
   git clone https://github.com/arxel2468/Go-Projects.git
   cd Go-Projects/Beginner/todo-app

Initialize Go Module

   ```bash
   go mod init github.com/arxel2468/Go-Projects/Beginner/todo-app
Run the Application

  ```bash
  go run main.go

Open in Browser
Visit http://localhost:8080

📝 Usage
Add Task: Enter a task in the input box and click "Add Task".
Delete Task: Click "Delete" next to a task to remove it.

📦 Dependencies
Standard Go libraries: net/http, html/template, encoding/json, sync