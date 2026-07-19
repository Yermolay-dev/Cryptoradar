package handlers

import (
	"crypto-radar/internal/models"
	"encoding/json"
	"log/slog"
	"net/http"
)

func HandleAlert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req models.AlertRequest
	var res models.AlertResponse
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	slog.Info("Created alert to money", "money", req.Coin)
	res.Message = "Your JSON succesffully received!"
	res.Status = "Succes!"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(res)
}
