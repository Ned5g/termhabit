package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Habit struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type HabitData struct {
	Habits      []Habit `json:"habits"`
	LastUpdated string  `json:"last_updated"`
}

func InitHabitFile(appName string) string {
	dataDir := filepath.Join(os.Getenv("HOME"), ".local", "share", appName)
	err := os.MkdirAll(dataDir, 0o755)
	if err != nil {
		fmt.Printf("failed to create dir: %v", err)
	}

	return filepath.Join(dataDir, "habits.json")
}

func ReadHabitFile(appName string) HabitData {
	file := filepath.Join(os.Getenv("HOME"), ".local", "share", appName, "habits.json")

	op, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Failed to read habit file: %v", err)
	}

	var habitStr HabitData

	err = json.Unmarshal(op, &habitStr)
	if err != nil {
		fmt.Printf("failed to unmarshal habit file: %v", err)
	}

	return habitStr
}

func WriteHabitFile(appName string, data HabitData) {
	file := filepath.Join(os.Getenv("HOME"), ".local", "share", appName, "habits.json")

	b, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(file, b, 0o644)
}
