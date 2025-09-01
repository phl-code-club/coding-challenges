package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func createTables(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS teams (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL
	);
	CREATE TABLE IF NOT EXISTS members (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		team_id INTEGER NOT NULL,
		FOREIGN KEY(team_id) REFERENCES teams(id)
	);
	CREATE TABLE IF NOT EXISTS questions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		part_1_description TEXT NOT NULL,
		part_1_answer TEXT NOT NULL,
		part_2_description TEXT NOT NULL,
		part_2_answer TEXT NOT NULL
	);
	CREATE TABLE IF NOT EXISTS answers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		team_id INTEGER NOT NULL,
		question_id INTEGER NOT NULL,
		answer TEXT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		return fmt.Errorf("unable to create tables: %e", err)
	}
	return nil
}

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "file:test.db")
	if err != nil {
		return nil, fmt.Errorf("unable to open db: %e", err)
	}
	err = createTables(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
