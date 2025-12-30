package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}
func Error(w http.ResponseWriter, statusCode int, errMsg string) {
	JSON(w, statusCode, map[string]string{"error": errMsg})
}

func Success(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusOK, data)
}

func Created(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusCreated, data)
}

func BadRequest(w http.ResponseWriter, errMsg string) {
	Error(w, http.StatusBadRequest, errMsg)
}

func NotFound(w http.ResponseWriter, errMsg string) {
	Error(w, http.StatusNotFound, errMsg)
}

func InternalServerError(w http.ResponseWriter, errMsg string) {
	Error(w, http.StatusInternalServerError, errMsg)
}

func Unauthorized(w http.ResponseWriter, errMsg string) {
	Error(w, http.StatusUnauthorized, errMsg)
}

func Forbidden(w http.ResponseWriter, errMsg string) {
	Error(w, http.StatusForbidden, errMsg)
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
