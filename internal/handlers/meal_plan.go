package handlers

import (
	"encoding/json"
	"log/slog"
	"meal-plan/internal/models"
	"net/http"
)

func (h *MealPlanHandler) GetWeeklyPlan(w http.ResponseWriter, r *http.Request) {
	weeklyPlan, err := h.Service.GetWeeklyPlan()
	if err != nil {
		http.Error(w, "Error getting weekly plan", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weeklyPlan)
}

func (h *MealPlanHandler) SaveWeeklyPlan(w http.ResponseWriter, r *http.Request) {
	var plan models.PlanItem
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}
	err := h.Service.SaveWeeklyPlan(plan)
	if err != nil {
		slog.Error("Error saving weekly plan: %v", err)
		http.Error(w, "Error saving weekly plan", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
