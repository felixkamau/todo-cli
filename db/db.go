package db

import (
	"database/sql"
	"log"

	"github.com/felixkamau/todo-cli/types"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB(dbname string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbname)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func CreateTable(db *sql.DB) {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			status TEXT NOT NULL,
			done BOOLEAN NOT NULL DEFAULT 0
		);
	`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal("Failed to prepare to create table statement:", err)
	}

	defer statement.Close()

	_, err = statement.Exec()
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

func InsertTasks(db *sql.DB, task types.Task) {
	insertSQl := `
		INSERT INTO tasks 
		(name, status, done) VALUES (?, ?, ?)
	`
	statement, err := db.Prepare(insertSQl)
	if err != nil {
		log.Fatal("Failed to prepare insert statement:", err)
	}

	defer statement.Close()
	_, err = statement.Exec(task.Name, task.Status, task.Done)

	if err != nil {
		log.Fatal("Failed to insert task:", err)
	}
}

func GetAllTasks(db *sql.DB) ([]types.Task, error) {
	rows, err := db.Query("SELECT id, name, status, done FROM tasks")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []types.Task

	for rows.Next() {
		var t types.Task
		err := rows.Scan(&t.ID, &t.Name, &t.Status, &t.Done)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// TODOS
// Implement
// MarkDone()
// DeleteTask()

func MarkDone(db *sql.DB, taskid int) error {
	updateSQL := `
		UPDATE tasks 
		SET done = 1 
		WHERE id = ?
	`
	statement, err := db.Prepare(updateSQL)
	if err != nil {
		log.Fatal("Failed to prepare update statement:", err)
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(taskid)
	if err != nil {
		log.Fatal("Failed to mark task as done:", err)
		return err
	}
	return nil
}

func DeleteTask(db *sql.DB, taskid int) error {
	deleteSql := `
		DELETE FROM tasks
		WHERE id = ?
	`
	statement, err := db.Prepare(deleteSql)
	if err != nil {
		log.Fatalf("Failed to prepare delete statement: %v", err)
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(taskid)
	if err != nil {
		log.Fatalf("Failed to delete task with id %d: %v", taskid, err)
	}
	return nil
}
