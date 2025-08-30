package handler

import (
	"encoding/json"
	"net/http"

	"github.com/justinbather/soundstash/server/internal/core/ports"
	"github.com/justinbather/soundstash/server/internal/core/types"
)

type httpHandler struct {
	downloader ports.Downloader
}

func New(downloader ports.Downloader) ports.HttpHandler {
	return &httpHandler{downloader}
}

func (*httpHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, types.HealthCheckResponse{Message: "ok"})
}

func (*httpHandler) Download(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, types.HealthCheckResponse{Message: "ok"})
}

func jsonResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
