package db

import (
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// types
var DB *sqlx.DB

func InitializeDB() error {
	dbFileName := "db/database.db"

	if err := os.MkdirAll(filepath.Dir(dbFileName), os.ModePerm); err != nil {
		return err
	}

	if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
		file, err := os.Create(dbFileName)
		if err != nil {
			return err
		}
		file.Close()
	}

	var err error
	DB, err = sqlx.Connect("sqlite3", dbFileName)
	if err != nil {
		return err
	}

	// Criação da tabela de usuários
	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		email TEXT UNIQUE,
		password TEXT
	)
	`)
	if err != nil {
		return err
	}

	// Criação da tabela de postagens
	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		author_id INTEGER,
		title TEXT,
		content TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)
	`)
	if err != nil {
		return err
	}

	return nil
}
