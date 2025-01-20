# CLI Task Manager

A simple command-line tool to manage daily tasks.

## Features
- Add new tasks.
- List all tasks.
- Mark tasks as done.
- Delete tasks.

## Usage
1. Build the project:
   ```bash
   go build main.go
2. Run commands:

Add a task:
   ```bash
   ./main add "Buy groceries"
List all tasks:
   ```bash
   ./main list
Mark a task as done:
   ```bash
   ./main done <task-id>
Delete a task:
   ```bash
   ./main delete <task-id>

3. Tasks are stored in tasks.json in the current directory.