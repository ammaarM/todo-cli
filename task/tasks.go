package task

import (
	"encoding/json"
	"os"
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

func LoadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var Tasks []Task
	if err := json.Unmarshal(data, &Tasks); err != nil {
		return nil, err
	}

	return Tasks, nil
}

func SaveTasks(filename string, Tasks []Task) error {
	data, err := json.MarshalIndent(Tasks, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
