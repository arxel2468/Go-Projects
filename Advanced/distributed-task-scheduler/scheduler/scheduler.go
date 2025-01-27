package scheduler

import (
	"database/sql"
	"log"
	"sync"

	_ "modernc.org/sqlite"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Status      string `json:"status"`
}

type Scheduler struct {
	DB *sql.DB
	mu sync.Mutex
}

func NewScheduler() *Scheduler {
	db, err := sql.Open("sqlite", "./tasks.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize the tasks table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		description TEXT,
		priority INTEGER,
		status TEXT
	)`)
	if err != nil {
		log.Fatalf("Failed to create tasks table: %v", err)
	}

	return &Scheduler{DB: db}
}

func (s *Scheduler) AddTask(task Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.Status = "Pending"
	_, err := s.DB.Exec("INSERT INTO tasks (id, description, priority, status) VALUES (?, ?, ?, ?)",
		task.ID, task.Description, task.Priority, task.Status)
	return err
}

func (s *Scheduler) GetAllTasks() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	rows, err := s.DB.Query("SELECT id, description, priority, status FROM tasks")
	if err != nil {
		log.Printf("Failed to fetch tasks: %v", err)
		return nil
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Description, &task.Priority, &task.Status); err == nil {
			tasks = append(tasks, task)
		}
	}
	return tasks
}
