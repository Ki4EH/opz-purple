package database

import (
	"database/sql"
	"fmt"
	"github.com/Ki4EH/opz-purple/internal/logger"
	_ "github.com/lib/pq"
)

func ConnectToDB(host string, port int, user, passw, name string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, passw, name)

	db, err := sql.Open("postgres", connStr)
	lg, _ := logger.NewFileLogger()
	if err != nil {
		lg.Error(fmt.Sprintf("Error connect to db %v", err))
		return nil, err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
