package database

import "database/sql"

// Uncomment if actually using SQLite, the first build will be fairly slow
// import _ "github.com/mattn/go-sqlite3"

type sqliteDatabase struct {
	Db *sql.DB
}

func NewSQLite(filename string) (Database, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	return &sqliteDatabase{db}, nil
}
