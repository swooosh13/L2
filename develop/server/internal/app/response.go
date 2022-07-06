package app

import (
	"encoding/json"
	"net/http"
)

type ResultResponse struct {
	Result interface{} `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func JsonResponse(w http.ResponseWriter, err bool, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err {
		errResp := ErrorResponse{data.(string)}
		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			http.Error(w, "json encoding error", http.StatusInternalServerError)
		}
		return
	}

	resp := ResultResponse{data}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "json encoding error", http.StatusInternalServerError)
	}
}
