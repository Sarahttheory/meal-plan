package handlers

import (
	"encoding/json"
	"meal-plan/internal/models"
	"net/http"
)

func (h *MealPlanHandler) GetDishes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	dishes, err := h.service.GetDishes(ctx)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Error getting dishes")
		return
	}
	h.respondWithJson(w, http.StatusOK, dishes)
}

func (h *MealPlanHandler) SaveDish(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var dish models.CreateDishInput
	if err := json.NewDecoder(r.Body).Decode(&dish); err != nil {
		h.respondWithJson(w, http.StatusBadRequest, "Error parsing request body")
		return
	}

	if err := h.validator.Struct(dish); err != nil {
		h.respondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err := h.service.SaveDish(ctx, dish)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Could not save dish")
		return
	}
	h.respondWithJson(w, http.StatusCreated, map[string]string{"dish": dish.Name})
}

func (h *MealPlanHandler) GetIngredients(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ingredients, err := h.service.GetIngredients(ctx)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Error get ingredients")
		return
	}
	h.respondWithJson(w, http.StatusOK, ingredients)
}
