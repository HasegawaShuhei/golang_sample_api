package config

import (
	"database/sql"
	"fmt"
	"log"
)

func OpenDB(driver, dsn string) *sql.DB {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal("OpenDB failed:", err)
	}

	fmt.Println("db connected!!")
	return db
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal("CloseDB failed:", err)
	}
}
