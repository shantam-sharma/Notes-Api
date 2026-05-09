package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	//prepare db manager
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}
	//connecction attempt
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL")

	return db, nil
}
