# 📝 Jcodes Task List Web App (Go Edition)

Welcome to the **Jcodes Task List Web App** – a simple yet functional to-do list web application built with Go's built-in `net/http` package. This app allows users to view, add, and manage a list of anime tasks right from their browser.

---

## 🌐 Features

- 🔸 **Homepage**: Clean UI with navigation buttons.
- 📃 **Show Tasks**: Displays the current list of tasks dynamically.
- ➕ **Add Task**: Users can submit a new task using a form.
- 🔁 **Redirect Handling**: Submitting a task redirects back to the task list with a `GET` request.
- 🖥️ **Console Support**: Built-in terminal methods to test task logic outside the browser.

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Server**: Built-in `net/http` web server
- **Frontend**: HTML (served via Go templates/strings)
- **Port**: `localhost:8080`

---

## 📂 Project Structure

```go
main.go        // All logic and HTTP handling in a single file

🚀 How to Run
Install Go from https://golang.org/dl/

Clone the repository
git clone https://github.com/yourusername/jcodes-tasklist-go.git
cd jcodes-tasklist-go

Run
go run main.go

Visit the app in your browser
http://localhost:8080

🧠 How It Works
/ — Homepage
Displays a welcome message and buttons to access task features.

/show-tasks — View Tasks
Lists all current tasks from the global taskItems array.

/add-task — Add Task Form
Displays a form where users can enter a new task.

/submit-task — Form Handling
Accepts POST requests, adds the new task to the slice, and redirects using http.StatusSeeOther.

📦 Additional Functions
Though not used in the web interface, you can test the app in the terminal by uncommenting the terminal section in main():

// printTasks(taskItems)
// taskItems = addTask(taskItems, "Watch Naruto")
// printTasks(taskItems)

🧑‍💻 Author
Jcodes
🎓 Beginner-friendly Go project with web-based interaction

✅ Future Ideas
Save tasks persistently using files or a database

Mark tasks as completed

Delete tasks

Style frontend with CSS or integrate Go templates

📃 License
This project is open-source and available under the MIT License.
