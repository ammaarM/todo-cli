package task

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	Completed     bool      `json:"completed"`
	DateStarted   time.Time `json:"date_started"`
	DateCompleted time.Time `json:"date_completed"`
}

func LoadTasks() ([]Task, string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, "", err
	}

	dataDir := filepath.Join(home, ".todo-cli")
	filename := filepath.Join(dataDir, "tasks.json")

	// Ensure directory exists
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if mkErr := os.MkdirAll(dataDir, 0755); mkErr != nil {
			return nil, "", mkErr
		}
	}

	// If file doesn’t exist, create it with empty JSON
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		empty := []Task{}
		data, _ := json.MarshalIndent(empty, "", "  ")
		if writeErr := os.WriteFile(filename, data, 0644); writeErr != nil {
			return nil, "", writeErr
		}
	}

	// Read and load tasks
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, "", err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, "", err
	}

	return tasks, filename, nil
}

func SaveTasks(filename string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
