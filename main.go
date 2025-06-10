package main

import (
	"fmt"
	"log"
	"os"

	"github.com/felixkamau/todo-cli/db"
	"github.com/felixkamau/todo-cli/types"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  Add a task: go run main.go add <task_name> <status> [done]")
		fmt.Println("  List tasks: go run main.go list")
		os.Exit(1)
	}

	command := os.Args[1]
	dbname := "./tasks.db"
	sqlite3db, err := db.ConnectDB(dbname)
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}
	defer sqlite3db.Close()

	// Always ensure table exists
	db.CreateTable(sqlite3db)

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go add <task_name> <status> [done]")
			return
		}
		taskName := os.Args[2]
		status := os.Args[3]
		done := len(os.Args) > 4 && os.Args[4] == "done"
		if done {
			status = "completed"
		}
		task := types.Task{
			Name:   taskName,
			Status: status,
			Done:   done,
		}
		db.InsertTasks(sqlite3db, task)
		fmt.Println("Task added successfully")

	case "list":
		tasks, err := db.GetAllTasks(sqlite3db)
		if err != nil {
			log.Fatal("Error getting tasks:", err)
		}
		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for _, task := range tasks {
			doneStatus := "❌"
			if task.Done {
				doneStatus = "✅"
			}
			fmt.Printf("[%d] %s - %s %s\n", task.ID, task.Name, task.Status, doneStatus)
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
