package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

const CreateTableQuery = `
CREATE TABLE IF NOT EXISTS responses (
	serial_num INTEGER,
	date_time TEXT,
	temp1 REAL,
	temp2 REAL,
	height REAL,
	latitude REAL,
	logitude REAL,
	status NTEGER
)
`

func Init(sqliteDatabase string) (err error) {
	if sqliteDatabase == "" {
		return errors.New("sqlite database name is not provided")
	}

	DB, err = sql.Open("sqlite3", sqliteDatabase)
	if err != nil {
		return err
	}

	_, err = DB.Exec(CreateTableQuery)
	if err != nil {
		return err
	}

	fmt.Printf("Database \"%s\" is set\n", sqliteDatabase)
	return nil
}
