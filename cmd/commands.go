package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"todo-cli/task"
)

func validateTaskID(id int, tasks []task.Task) error {
	if id < 1 || id > len(tasks) {
		return fmt.Errorf("task ID is not valid")
	}
	return nil
}

func Execute() error {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: todoctl [command] [arguments]. Use --help or -h for more information.")
		return nil
	}

	tasks, filename, err := task.LoadTasks()
	if err != nil {
		log.Fatal(err)
	}

	switch args[0] {
	case "list":
		fmt.Println("--- Your Tasks ---")
		if len(args) > 1 && args[1] == "completed" {
			fmt.Println("Completed Tasks:")
			for _, t := range tasks {
				if t.Completed {
					status := "x"
					fmt.Printf("[%s] %d: %s (Type: %s)\n", status, t.ID, t.Name, t.Type)
				}
			}
			return nil
		} else if len(args) > 1 && args[1] == "uncompleted" {
			fmt.Println("Uncompleted Tasks:")
			for _, t := range tasks {
				if !t.Completed {
					status := " "
					fmt.Printf("[%s] %d: %s (Type: %s)\n", status, t.ID, t.Name, t.Type)
				}
			}
			return nil
		} else {
			for _, t := range tasks {
				status := " "
				if t.Completed {
					status = "x"
				}
				fmt.Printf("[%s] %d: %s (Type: %s)\n", status, t.ID, t.Name, t.Type)
			}
		}

	case "get":
		if len(args) < 2 {
			fmt.Println("Usage: todo-cli add <task id>")
			return nil
		}

		id, _ := strconv.Atoi(args[1])
		if err := validateTaskID(id, tasks); err != nil {
			fmt.Println(err)
			return nil
		}

		for i, t := range tasks {
			if t.ID == id {
				tasks[i].Completed = true
				tasks[i].DateCompleted = time.Now()

				status := " "
				if t.Completed {
					status = "x"
				}
				fmt.Printf("[%s] \n %d: \n %s \n Type: %s \n Date Started: %s \n Date Completed: %s \n", status, t.ID, t.Name, t.Type, t.DateStarted, t.DateCompleted)
			}
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
		task.SaveTasks(filename, tasks)
		fmt.Println("Added:", name)

	case "complete":
		if len(args) < 2 {
			fmt.Println("Usage: todo-cli complete <task id>")
			return nil
		}
		id, _ := strconv.Atoi(args[1])
		if err := validateTaskID(id, tasks); err != nil {
			fmt.Println(err)
			return nil
		}

		for i, t := range tasks {
			if t.ID == id {
				tasks[i].Completed = true
				tasks[i].DateCompleted = time.Now()
			}
		}
		task.SaveTasks(filename, tasks)
		fmt.Println("Completed task ID:", id)

	case "uncomplete":
		if len(args) < 2 {
			fmt.Println("Usage: todo-cli complete <task id>")
			return nil
		}
		id, _ := strconv.Atoi(args[1])
		if err := validateTaskID(id, tasks); err != nil {
			fmt.Println(err)
			return nil
		}

		for i, t := range tasks {
			if t.ID == id {
				tasks[i].Completed = false
			}
		}
		task.SaveTasks(filename, tasks)
		fmt.Println("Uncompleted task ID:", id)

	case "delete":
		if len(args) < 2 {
			fmt.Println("Usage: todo-cli delete <task ID>")
			return nil
		}
		id, _ := strconv.Atoi(args[1])
		if err := validateTaskID(id, tasks); err != nil {
			fmt.Println(err)
			return nil
		}

		for i := range tasks {
			if tasks[i].ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				break
			}
		}
		task.SaveTasks(filename, tasks)
		fmt.Println("Deleted:", id)
	case "--help", "-h":
		fmt.Println("Usage: todoctl [command] [arguments]")
		fmt.Println("Commands:")
		fmt.Println("  list <optional: completed/uncompleted>	Display all tasks")
		fmt.Println("  add <task name>        					Add a new task")
		fmt.Println("  complete <task id>                       Mark a task as completed")
		fmt.Println("  uncomplete <task id>                     Mark a task as not completed")
		fmt.Println("  delete <task id>                         Delete a task")
		fmt.Println("  --help, -h                               Show this help message")
	default:
		fmt.Println("todoctl:", args[0], "is not a command. See 'todoctl --help'")
	}

	return nil
}
