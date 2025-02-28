package main

import (
    "encoding/json"
    "net/http"
    "time"
)

type TimeResponse struct {
    Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
    response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/time", timeHandler)
    http.ListenAndServe(":8795", nil)
}
