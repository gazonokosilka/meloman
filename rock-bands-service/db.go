package main

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite", "artists.db")
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	createTablesQuery := `
		CREATE TABLE IF NOT EXISTS artists (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			born TEXT NOT NULL,
			genre TEXT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS albums (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			year TEXT NOT NULL,
			artist_id TEXT NOT NULL,
			FOREIGN KEY (artist_id) REFERENCES artists (id)
		);
	`
	_, err = db.Exec(createTablesQuery)
	if err != nil {
		return fmt.Errorf("ошибка при создании таблиц: %w", err)
	}

	return nil
}
