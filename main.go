package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	UTC  string `json:"utc_time"`
	Kyiv string `json:"kyiv_time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().UTC()

	kyivLocation, err := time.LoadLocation("Europe/Kyiv")
	if err != nil {
		http.Error(w, "Ошибка загрузки часового пояса", http.StatusInternalServerError)
		return
	}

	kyivTime := now.In(kyivLocation)

	response := TimeResponse{
		UTC:  now.Format(time.RFC3339),
		Kyiv: kyivTime.Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8795", nil)
}
