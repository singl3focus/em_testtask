package handler

import (
	"net/http"
	"log/slog"
)

type Service interface {

}

type Handler struct {
	logger *slog.Logger
	service Service
}

func NewHandler(service Service, logger *slog.Logger) *Handler {
	return &Handler{
		logger: logger,
		service: service,
	}
}

func (h *Handler) Healthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}