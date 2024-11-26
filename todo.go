package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Task structure to represent a to-do item
type Task struct {
	ID     int
	Title  string
	IsDone bool
}

var tasks []Task
var nextID = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- To-Do List ---")
		fmt.Println("1. Add a Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete a Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addTask(scanner)
		case "2":
			viewTasks()
		case "3":
			markTaskDone(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			fmt.Println("Exiting... Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task title: ")
	scanner.Scan()
	title := scanner.Text()

	task := Task{
		ID:     nextID,
		Title:  title,
		IsDone: false,
	}
	tasks = append(tasks, task)
	nextID++

	fmt.Println("Task added successfully!")
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("\n--- Your Tasks ---")
	for _, task := range tasks {
		status := "Pending"
		if task.IsDone {
			status = "Done"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, task.Title, status)
	}
}

func markTaskDone(scanner *bufio.Scanner) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to mark as done.")
		return
	}

	fmt.Print("Enter task ID to mark as done: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].IsDone = true
			fmt.Println("Task marked as done!")
			return
		}
	}

	fmt.Println("Task not found.")
}

func deleteTask(scanner *bufio.Scanner) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to delete.")
		return
	}

	fmt.Print("Enter task ID to delete: ")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			return
		}
	}

	fmt.Println("Task not found.")
}
