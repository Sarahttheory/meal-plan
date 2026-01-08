package handlers

import (
	"encoding/json"
	"meal-plan/internal/service"
	"net/http"
)

type MealPlanHandler struct {
	Service *service.MealPlanService
}

func (h *MealPlanHandler) GetAllDishes(w http.ResponseWriter, r *http.Request) {
	dishes, err := h.Service.GetDishes()
	if err != nil {
		http.Error(w, "Error getting dishes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dishes)
}
