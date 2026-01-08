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
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		panic(err)
	}

	file := filepath.Join(dataDir, "habits.json")

	if _, err := os.Stat(file); os.IsNotExist(err) {
		defaultData := HabitData{
			Habits:      []Habit{},
			LastUpdated: "",
		}

		b, _ := json.MarshalIndent(defaultData, "", " ")
		_ = os.WriteFile(file, b, 0o644)
	}

	return file
}

func ReadHabitFile(appName string) HabitData {
	file := filepath.Join(os.Getenv("HOME"), ".local", "share", appName, "habits.json")

	op, err := os.ReadFile(file)
	if err != nil {
		InitHabitFile(appName)
	}

	var habitStr HabitData

	err = json.Unmarshal(op, &habitStr)
	if err != nil {
		fmt.Printf("Failed to unmarshal habit file: %v", err)
	}

	return habitStr
}

func WriteHabitFile(appName string, data HabitData) error {
	dataDir := filepath.Join(os.Getenv("HOME"), ".local", "share", appName)
	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return fmt.Errorf("failed to create data dir: %w", err)
	}

	file := filepath.Join(dataDir, "habits.json")

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal habit data: %w", err)
	}

	if err := os.WriteFile(file, b, 0o644); err != nil {
		return fmt.Errorf("failed to write habit file: %w", err)
	}

	return nil
}
