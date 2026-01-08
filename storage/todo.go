package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Todo struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type TodoData struct {
	Todos []Todo `json:"todos"`
}

func InitTodoFile(appName string) string {
	dataDir := filepath.Join(os.Getenv("HOME"), ".local", "share", appName)
	err := os.MkdirAll(dataDir, 0o755)
	if err != nil {
		fmt.Printf("failed to create dir: %v", err)
	}

	return filepath.Join(dataDir, "todo.json")
}

func ReadTodoFile(appName string) TodoData {
	file := filepath.Join(os.Getenv("HOME"), ".local", "share", appName, "todo.json")

	op, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("failed to read habit file: %v", err)
	}

	var todoStr TodoData

	err = json.Unmarshal(op, &todoStr)
	if err != nil {
		fmt.Printf("failed to unmarshal habit file: %v", err)
	}

	return todoStr
}

func WriteTodoFile(appName string, data TodoData) {
	file := filepath.Join(os.Getenv("HOME"), ".local", "share", appName, "todo.json")

	b, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile(file, b, 0o644)
}
