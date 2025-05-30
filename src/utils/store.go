package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"todo/src/models"
)

func SaveTasks(tasks []models.Task) {
	// parse json file
	fileByte, err := json.MarshalIndent(tasks, "", "  ")
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
