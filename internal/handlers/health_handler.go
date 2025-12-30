package handlers

import (
	"net/http"

	"github.com/tobiasaagaard/localgift-api/pkg/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response.Success(w, map[string]string{
		"service": "LocalGift API",
		"status":  "healthy",
	})
}
