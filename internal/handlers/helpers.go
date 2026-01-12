package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *MealPlanHandler) respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (h *MealPlanHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	h.respondWithJson(w, code, map[string]string{"error": message})
}
