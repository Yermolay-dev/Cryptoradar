package main

import (
	"crypto-radar/internal/handlers"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/alert", handlers.HandleAlert)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Error("Received the error", "Error", err.Error())
	}
}
