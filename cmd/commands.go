package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"todo-cli/task"
)

func Execute() error {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: todo-cli [list|add|complete|delete]")
		return nil
	}

	tasks, err := task.LoadTasks("tasks.json")
	if err != nil {
		return err
	}

	switch args[0] {
	case "list":
		for _, t := range tasks {
			status := " "
			if t.Completed {
				status = "x"
			}
			fmt.Printf("[%s] %d: %s (Type: %s)\n", status, t.ID, t.Name, t.Type)
		}

	case "add":
		if len(args) < 2 {
			fmt.Println("Usage: todo-cli add <task name>")
			return nil
		}
		name := args[1]
		newTask := task.Task{
			ID:          len(tasks) + 1,
			Name:        name,
			DateStarted: time.Now(),
		}
		tasks = append(tasks, newTask)
		task.SaveTasks("tasks.json", tasks)
		fmt.Println("Added:", name)

	case "complete":
		if len(args) < 2 {
			fmt.Println("Usage: todo-cli complete <task id>")
			return nil
		}
		id, _ := strconv.Atoi(args[1])
		for i, t := range tasks {
			if t.ID == id {
				tasks[i].Completed = true
				tasks[i].DateCompleted = time.Now()
			}
		}
		task.SaveTasks("tasks.json", tasks)
		fmt.Println("Completed task ID:", id)

	case "delete":
		if len(args) < 2 {
			fmt.Println("Usage: todo-cli delete <task ID>")
			return nil
		}
		id, _ := strconv.Atoi(args[1])
		for i := range tasks {
			if tasks[i].ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}
		task.SaveTasks("tasks.json", tasks)
		fmt.Println("Deleted:", id)

	default:
		fmt.Println("Unknown command:", args[0])
	}

	return nil
}
