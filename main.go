package main

import (
	"fmt"
	"log"
	"os"

	"github.com/felixkamau/todo-cli/db"
	"github.com/felixkamau/todo-cli/types"
)

func main() {
	// Check args
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <task_name> <status> [done]")
		fmt.Println("Example: go run main.go \"Test code\" \"Pending\"")
		fmt.Println("Example: go run main.go \"Fix bug\" \"in-progress\" done")
		os.Exit(1) // Exit if not enough arguments
	}

	// Parse cli args (this should be OUTSIDE the if block)
	taskName := os.Args[1]
	status := os.Args[2]
	done := false

	// Check if "done" flag is provided
	if len(os.Args) > 3 && os.Args[3] == "done" {
		done = true
		status = "completed"
	}

	// create task
	task := types.Task{
		Name:   taskName,
		Status: status,
		Done:   done,
	}

	dbname := "./tasks.db"
	sqlite3db, err := db.ConnectDB(dbname)
	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}

	defer sqlite3db.Close()

	// Create table
	db.CreateTable(sqlite3db)

	// Insert/add task
	db.InsertTasks(sqlite3db, task)

	fmt.Println("Task added successfully")
	fmt.Printf("Name: %s\n", task.Name)
	fmt.Printf("Status: %s\n", task.Status)
	fmt.Printf("Done: %t\n", task.Done)
}
