package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo/src/models"
	"todo/src/utils"
)

var taskData []models.Task

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

			newTask := models.Task{
				ID:        len(taskData) + 1,
				Name:      taskName,
				Completed: false,
			}
			taskData = append(taskData, newTask)

			utils.SaveTasks(taskData)
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
					utils.SaveTasks(taskData)
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
					utils.SaveTasks(taskData)
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
