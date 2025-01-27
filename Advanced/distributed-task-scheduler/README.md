# Distributed Task Scheduler

A simple yet extensible distributed task scheduler built with **Go** and **Gin Web Framework**. This project allows users to add, manage, and prioritize tasks with ease, providing a clean and user-friendly interface.

## Features
- Add tasks with a description, priority, and status.
- View a list of all tasks.
- Lightweight and efficient, designed for scalability.
- Modular structure for easy maintenance and extensibility.

## Getting Started

### Prerequisites
Ensure you have the following installed on your machine:
- [Go (1.20+)](https://go.dev/dl/)
- A web browser to access the UI

### Setup
1. Clone this repository:
   ```bash
   git clone https://github.com/yourusername/distributed-task-scheduler.git
   cd distributed-task-scheduler
2. Install dependencies:
    ```bash
    go mod tidy
3. Run the server
    ```bash
    go run main.go
4. Open your browser and navigate to:
    ```arduino
    http://localhost:8080
    

## Usage
1. Adding Tasks

- Enter the task details (ID, description, priority).
- Click Add Task to add the task to the list.

2. Viewing Tasks
- The task list is displayed below the form with their ID, description, priority, and status.

## Testing
To test the project, follow these steps:

1. Add various tasks through the UI and verify that they appear in the list.
2. Manually inspect the /tasks API endpoint:
- Open a tool like Postman or use curl:
    ```bash
    curl http://localhost:8080/tasks

## API Endpoints
- GET /tasks: Retrieve all tasks.
- POST /tasks: Add a new task.
Example payload:
    ```json
    {
        "id": "1",
        "description": "Implement feature X",
        "priority": 10,
        "status": "Pending"
    }

## Future Enhancements
- Add a database for persistent storage.
- Implement task editing and deletion features.
- Introduce distributed execution for tasks across multiple nodes.