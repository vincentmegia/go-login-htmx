package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase(file string) (*sql.DB, error) {
	db, error := sql.Open("sqlite3", "./customers.db")
	if error != nil {
		fmt.Println("Cannot open database", error)
		return nil, error
	}
	fmt.Println("Opened database success")
	return db, nil
}
