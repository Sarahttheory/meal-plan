package handlers

import (
	"encoding/json"
	"log/slog"
	"meal-plan/internal/models"
	"net/http"
)

func (h *MealPlanHandler) GetDishes(w http.ResponseWriter, r *http.Request) {
	dishes, err := h.Service.GetDishes()
	if err != nil {
		http.Error(w, "Error getting dishes", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dishes)
}

func (h *MealPlanHandler) SaveDish(w http.ResponseWriter, r *http.Request) {
	var dish models.Dish
	if err := json.NewDecoder(r.Body).Decode(&dish); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}
	err := h.Service.SaveDish(dish)
	if err != nil {
		slog.Error("Error saving dish: %v", err)
		http.Error(w, "Error saving dish", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *MealPlanHandler) GetIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.Service.GetIngredients()
	if err != nil {
		http.Error(w, "Error getting ingredients", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredients)
}
