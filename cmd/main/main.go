package main

import (
	"fmt"
	"github.com/Ki4EH/opz-purple/env"
	"github.com/Ki4EH/opz-purple/internal/logger"
	"github.com/Ki4EH/opz-purple/pkg/database"
	_ "github.com/Ki4EH/opz-purple/pkg/database"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	log, _ := logger.NewFileLogger()

	db, err := database.ConnectToDB(env.Host, env.Port, env.User, env.Password, env.Dbname)

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	log.Info("[SERVER] Starting server")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Error(fmt.Sprintf("Error starting server %v", err))
	}
	defer log.Close()
}
