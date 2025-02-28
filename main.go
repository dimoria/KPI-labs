package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Структура JSON-ответа
type TimeResponse struct {
	UTC  string `json:"utc_time"`
	Kyiv string `json:"kyiv_time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем текущее UTC-время
	now := time.Now().UTC()

	// Загружаем часовой пояс Киева
	kyivLocation, err := time.LoadLocation("Europe/Kyiv")
	if err != nil {
		http.Error(w, "Ошибка загрузки часового пояса", http.StatusInternalServerError)
		return
	}

	// Конвертируем UTC в киевское время
	kyivTime := now.In(kyivLocation)

	// Создаём JSON-ответ
	response := TimeResponse{
		UTC:  now.Format(time.RFC3339),
		Kyiv: kyivTime.Format(time.RFC3339),
	}

	// Устанавливаем заголовок Content-Type для JSON
	w.Header().Set("Content-Type", "application/json")

	// Отправляем JSON-ответ
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8795", nil)
}

//new comment for revert
