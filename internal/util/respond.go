package util

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func Error(w http.ResponseWriter, err error) {
	if ae, ok := err.(AppError); ok {
		JSON(w, ae.Code, map[string]string{"error": ae.Message})
		return
	}
	JSON(w, http.StatusInternalServerError, map[string]string{"error": "internal error"})
}
