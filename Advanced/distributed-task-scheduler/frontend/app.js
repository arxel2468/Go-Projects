document.getElementById("taskForm").addEventListener("submit", async (e) => {
    e.preventDefault(); // Prevents page refresh

    const id = document.getElementById("id").value;
    const description = document.getElementById("description").value;
    const priority = parseInt(document.getElementById("priority").value);

    try {
        const response = await fetch("/tasks", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ id, description, priority }),
        });

        if (response.ok) {
            alert("Task added!");
            loadTasks();
        } else {
            alert("Failed to add task. Please try again.");
        }
    } catch (error) {
        console.error("Error adding task:", error);
        alert("An error occurred. Please check the console for details.");
    }
});

async function loadTasks() {
    try {
        const res = await fetch("/tasks");
        if (res.ok) {
            const tasks = await res.json();
            const taskList = document.getElementById("taskList");
            taskList.innerHTML = "";

            tasks.forEach(task => {
                const li = document.createElement("li");
                li.textContent = `${task.id}: ${task.description} (Priority: ${task.priority}, Status: ${task.status})`;
                taskList.appendChild(li);
            });
        } else {
            console.error("Failed to fetch tasks");
        }
    } catch (error) {
        console.error("Error fetching tasks:", error);
    }
}

loadTasks();
