package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h *MealPlanHandler) respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		slog.Error("failed to marshal json response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (h *MealPlanHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	h.respondWithJson(w, code, map[string]string{"error": message})
}
