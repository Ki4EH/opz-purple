package api

import (
	"encoding/json"
	"github.com/Ki4EH/opz-purple/internal/models"
	"github.com/Ki4EH/opz-purple/pkg/database"
	"net/http"
)

func AddPrice(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req models.RequestAddPrice
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()

	err := database.Connection.AddNewPrice(req)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ReturnPrice(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var req models.RequestPrice
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()
	//TODO: отправить ответ из (baseline, discount) матриц
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func SetupRoutes() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/user/", func(writer http.ResponseWriter, request *http.Request) {
		ReturnPrice(writer, request)
	})
	router.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		AddPrice(writer, request)
	})
	return router
}
