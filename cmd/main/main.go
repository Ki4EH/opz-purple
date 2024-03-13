package main

import (
	"fmt"
	"github.com/Ki4EH/opz-purple/internal/api"
	"github.com/Ki4EH/opz-purple/internal/config"
	"github.com/Ki4EH/opz-purple/internal/logger"
	_ "github.com/Ki4EH/opz-purple/pkg/database"
	"github.com/Ki4EH/opz-purple/pkg/treebase/discount"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	logging, _ := logger.NewFileLogger()
	//storage, err := database.NewStorage()
	//if err != nil {
	//	panic(err)
	//}
	//_ = storage

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      api.SetupRoutes(),
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	fmt.Println(discount.GetSegmentsByID(1000))
	fmt.Println(discount.GetSegmentsByID(4500))

	logging.Info("[SERVER] Starting server")

	if err := srv.ListenAndServe(); err != nil {
		logging.Error(fmt.Sprintf("Error starting server %v", err))
	}
	defer logging.Close()
}
