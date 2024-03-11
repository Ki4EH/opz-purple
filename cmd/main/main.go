package main

import (
	"fmt"
	"github.com/Ki4EH/opz-purple/internal/logger"
	"net/http"
)

func main() {
	log, _ := logger.NewFileLogger()
	log.Info("[SERVER] Starting server")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Error(fmt.Sprintf("Error starting server %v", err))
	}
	defer log.Close()
}
