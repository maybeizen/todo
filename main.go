package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var taskData []Task

func main() {
	fileByte, err := os.ReadFile("./data/tasks.json")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	err = json.Unmarshal(fileByte, &taskData)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose an option:\n1. Add a task\n2. List tasks\n3. Complete a task\n4. Delete a task\n5. Exit\n>> ")

	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	fmt.Println(input)

	switch input {

	case "1":
		{
			fmt.Println("Enter a task name: ")
			scanner.Scan()
			taskName := strings.TrimSpace(scanner.Text())
			taskData = append(taskData, Task{ID: len(taskData) + 1, Name: taskName, Completed: false})
			fmt.Println("Task added successfully!")
			break
		}
	case "2":
		{
			if len(taskData) == 0 {
				fmt.Println("No tasks found.")
			}

			fmt.Println("List of tasks:")
			for _, task := range taskData {
				fmt.Println(task.ID, task.Name, task.Completed)
			}
			break
		}

	case "3":
		{
			fmt.Println("Enter the task ID to complete:")
			scanner.Scan()
			taskID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid task ID. Must be a number.")
			}
			for i, task := range taskData {
				if task.ID == taskID {
					taskData[i].Completed = true
					break
				} else {
					fmt.Println("Task not found.")
				}
			}
			break
		}

	}

}
