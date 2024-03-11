package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectToDB(host string, port int, user, passw, name string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, passw, name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
