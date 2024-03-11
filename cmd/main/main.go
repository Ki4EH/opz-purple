package main

import (
	"fmt"
	"github.com/Ki4EH/opz-purple/pkg/database"

	//"github.com/Ki4EH/opz-purple/env"
	"github.com/Ki4EH/opz-purple/internal/config"
	"github.com/Ki4EH/opz-purple/internal/logger"
	_ "github.com/Ki4EH/opz-purple/pkg/database"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	log, _ := logger.NewFileLogger()

	_, err := database.ConnectToDB(cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.Name)

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
