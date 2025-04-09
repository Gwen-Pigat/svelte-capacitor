package initializers

import (
	"database/sql"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_URI"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
