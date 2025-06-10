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
