package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// task types
type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var taskData []Task

func main() {
	fileByte, err := os.ReadFile("./data/tasks.json")
	if err != nil {
		fmt.Println("❌ Error opening file:", err)
		return
	}

	// parse json file
	err = json.Unmarshal(fileByte, &taskData)
	if err != nil {
		fmt.Println("❌ Error parsing JSON:", err)
		return
	}

	// scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n====== TODO CLI ======")
		fmt.Println("1) Add a task")
		fmt.Println("2) List tasks")
		fmt.Println("3) Complete a task")
		fmt.Println("4) Delete a task")
		fmt.Println("5) Exit")
		fmt.Print(">> ")

		scanner.Scan()
		// trim whitespace from inputs
		input := strings.TrimSpace(scanner.Text())

		switch input {

		case "1":
			fmt.Print("\nEnter a task name: ")
			scanner.Scan()
			taskName := strings.TrimSpace(scanner.Text())

			newTask := Task{
				ID:        len(taskData) + 1,
				Name:      taskName,
				Completed: false,
			}
			taskData = append(taskData, newTask)

			saveTasks()
			fmt.Println("✅ Task added successfully!")

		case "2":
			if len(taskData) == 0 {
				fmt.Println("⚠️  No tasks found.")
				continue
			}

			fmt.Println("\n📋 List of tasks:")
			for _, task := range taskData {
				status := "[ ]"
				if task.Completed {
					status = "[x]"
				}
				fmt.Printf("%s %d: %s\n", status, task.ID, task.Name)
			}

		case "3":
			fmt.Print("\nEnter the task ID to mark as complete: ")
			scanner.Scan()
			taskID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("❌ Invalid task ID. Must be a number.")
				continue
			}

			found := false
			for i, task := range taskData {
				if task.ID == taskID {
					taskData[i].Completed = true
					saveTasks()
					fmt.Println("✅ Task marked as complete!")
					found = true
					break
				}
			}
			if !found {
				fmt.Println("❌ Task not found.")
			}

		case "4":
			fmt.Print("\nEnter the task ID to delete: ")
			scanner.Scan()
			taskID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("❌ Invalid task ID. Must be a number.")
				continue
			}

			found := false
			for i, task := range taskData {
				if task.ID == taskID {
					taskData = append(taskData[:i], taskData[i+1:]...)
					saveTasks()
					fmt.Println("✅ Task deleted successfully.")
					found = true
					break
				}
			}
			if !found {
				fmt.Println("❌ Task not found.")
			}

		case "5":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("⚠️  Invalid option. Please enter 1–5.")
		}
	}
}

// helper function to save tasks to disk
func saveTasks() {
	// parse json file
	fileByte, err := json.MarshalIndent(taskData, "", "  ")
	if err != nil {
		fmt.Println("❌ Error writing to file:", err)
		return
	}
	// write to disk
	err = os.WriteFile("./data/tasks.json", fileByte, 0644)
	if err != nil {
		fmt.Println("❌ Error writing to file:", err)
	}
}
