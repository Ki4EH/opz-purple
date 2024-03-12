package main

import (
	"fmt"
	"github.com/Ki4EH/opz-purple/internal/config"
	"github.com/Ki4EH/opz-purple/internal/logger"
	"github.com/Ki4EH/opz-purple/pkg/database"
	_ "github.com/Ki4EH/opz-purple/pkg/database"
	"github.com/Ki4EH/opz-purple/pkg/treebase/location"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	logging, _ := logger.NewFileLogger()
	_, err := database.ConnectToDB(cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.Name)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	tree := location.GetLocationsTree()
	tree.PrintTree(0)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      nil,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
	logging.Info("[SERVER] Starting server")

	if err := srv.ListenAndServe(); err != nil {
		logging.Error(fmt.Sprintf("Error starting server %v", err))
	}
	defer logging.Close()
}
