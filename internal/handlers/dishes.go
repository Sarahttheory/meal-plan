package handlers

import (
	"encoding/json"
	"meal-plan/internal/models"
	"net/http"
)

func (h *MealPlanHandler) GetDishes(w http.ResponseWriter, r *http.Request) {
	dishes, err := h.service.GetDishes()
	if err != nil {
		http.Error(w, "Error getting dishes", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dishes)
}

func (h *MealPlanHandler) SaveDish(w http.ResponseWriter, r *http.Request) {
	var dish models.CreateDishInput
	if err := json.NewDecoder(r.Body).Decode(&dish); err != nil {
		h.respondWithJson(w, http.StatusBadRequest, "Error parsing request body")
		return
	}

	if err := h.validator.Struct(dish); err != nil {
		h.respondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err := h.service.SaveDish(dish)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Could not save dish")
		return
	}
	h.respondWithJson(w, http.StatusCreated, map[string]string{"dish": dish.Name}) //todo заменить везде на красивые выводы
}

func (h *MealPlanHandler) GetIngredients(w http.ResponseWriter, r *http.Request) {
	ingredients, err := h.service.GetIngredients()
	if err != nil {
		http.Error(w, "Error getting ingredients", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ingredients)
}
