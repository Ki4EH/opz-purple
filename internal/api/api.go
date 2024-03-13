package api

import (
	"encoding/json"
	"github.com/Ki4EH/opz-purple/internal/models"
	"github.com/Ki4EH/opz-purple/pkg/database"
	"github.com/Ki4EH/opz-purple/pkg/treebase/discount"
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

func UpdatePrice(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req models.RequestAddPrice
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()

	database.Connection.UpdatePrice(req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// TODO: нам нужно сделать хендлер на увелмчение в процентаже всех локайи(категорий)
func UpdatePrices(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req models.RequestWithPercentage
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()

	if req.Price != 0 {
		database.Connection.UpdateManyPrices(req)
	} else if req.Percent != 0 {
		database.Connection.UpdateManyPricesWithPercentage(req)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

//TODO: сделать создание таблицы

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
	//TODO: сегмент получение
	seg, seg_slice := discount.GetSegmentByID(req.UserId)
	ans := database.SearchData(seg_slice, req)
	ans.UserSegmentId = seg
	w.Header().Set("Content-Type", "application/json")
	answer, _ := json.Marshal(ans)
	w.Write(answer)
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
	router.HandleFunc("/update", func(writer http.ResponseWriter, request *http.Request) {
		UpdatePrice(writer, request)
	})
	router.HandleFunc("/update/percentage", func(writer http.ResponseWriter, request *http.Request) {
		UpdatePrices(writer, request)
	})
	return router
}
