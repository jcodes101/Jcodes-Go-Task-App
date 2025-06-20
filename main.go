package main

import (
	"fmt"
	"net/http"
)

var task1 string = "Watch Solo Leveling"
var task2 string = "Watch Fairytale"
var task3 string = "watch One Piece"

var taskItems = []string{task1, task2, task3}

func main() {

	fmt.Println("|-| Welcome to Jcodes Task List App! |-|")
	fmt.Println("=========================================")
	fmt.Println()

	// GLOBAL VARIABLES FOR TERMINAL EXECUTION ==========================================
	// var task1 string = "Watch Solo Leveling"
	// var task2 string = "Watch Fairytale"
	// var task3 string = "watch One Piece"

	// var taskItems = []string{task1, task2, task3}

	// printTasks(taskItems)
	// fmt.Println()

	// fmt.Println("-Your updated task list-")
	// taskItems = addTask(taskItems, "Watch Naruto")

	// printTasks(taskItems)

	// ==================================================================================

	// HTTP SERVER SETUP

	// this is a built-in http function to handle requests made to the http server
	// for each endpoint to direct or redirect a user to a specific page or function
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.HandleFunc("/add-task", addTaskForm)
	http.HandleFunc("/submit-task", newTaskSubmission)

	// Local server to make application available to end-point users using port :8080
	http.ListenAndServe(":8080", nil)
}

// EXECUTED VIA HTTP REQUEST / BROWSER ==================================================

/*
GREETING FUNCTION

a (*http.Request) is being requested (r = request) from the user
and a (http.ResponseWriter) is being written (w = writer) to the user
*/
func helloUser(w http.ResponseWriter, r *http.Request) {
	// tells the browswer the type of content being sent to then be rendered as such
	w.Header().Set("Content-Type", "text/html")

	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
			<head>
				<title style="color: white">Welcome</title>
			</head>
			<body style="background-color:gray;">
				<h1 style="color:black">| -- Welcome to Jcodes Task List Web-App! -- |</h1>
				<button onclick="location.href='/show-tasks'">Show Tasks</button>
				<button onclick="location.href='/add-task'">Add Task</button>
			</body>
		</html>
	`)
}

/*
SHOW TASKS FUNCTION

upon entering the /show-tasks endpoint, the user will be shown
a list of tasks that are in their task list
*/
func showTasks(w http.ResponseWriter, r *http.Request) {
	var response string = "Here are your tasks:\n"
	fmt.Fprintf(w, response)

	for i, task := range taskItems {
		fmt.Fprintln(w, i+1, ".", task)
	}
}

/*
ADD TASK FORM FUNCTION

this is the form which allows a user to add a task
to their existsing task list.
*/
func addTaskForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
			<head>
				<title style="color:black">Add Task</title>
			</head>
			
			<body style="background-color:gray;">
				<h2>Add a new task</h2>
				<form method="POST" action="/submit-task">
					<input type="text" name="task-input" placeholder="List a new task here..." required>
					<input type="submit" value="Add Task">
				</form>
				<br>
				<a href="/show-tasks">View Tasks</a>
			</body>
		</html>
	`)
}

/*
NEW TASK SUBMISSION FUNCTION

this function handles the submission of a new task
when the user submits a task via the form
*/
func newTaskSubmission(w http.ResponseWriter, r *http.Request) {
	// as long as the request method is POST, the form will be processed
	if r.Method == http.MethodPost {
		// a variable task is assigned the value of the input field
		task := r.FormValue("task-input")
		if task != "" {
			taskItems = append(taskItems, task)
			// after form submission, the user is automatically redirected to the
			// /show-tasks endpoint; the (StatusSeeOther) tells the broswer to then
			// use the GET method to retrieve the /show-tasks page
			http.Redirect(w, r, "/show-tasks", http.StatusSeeOther)

			return
		}
	}
	http.Error(w, "Invalid task submission", http.StatusBadRequest)
}

// EXECUTED VIA TERMINAL ================================================================
func printTasks(taskItems []string) {
	fmt.Println("Animes to watch:")
	for i, task := range taskItems {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	updatedTaskItems := append(taskItems, newTask)
	return updatedTaskItems
}
