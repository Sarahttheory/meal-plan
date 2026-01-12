package handlers

import (
	"encoding/json"
	"meal-plan/internal/models"
	"net/http"
)

func (h *MealPlanHandler) GetWeeklyPlan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	weeklyPlan, err := h.service.GetWeeklyPlan(ctx)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Error getting weekly plan")
		return
	}
	h.respondWithJson(w, http.StatusOK, weeklyPlan)
}

func (h *MealPlanHandler) SaveItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var item models.PlanItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Error parsing request body")
		return
	}
	if err := h.validator.Struct(item); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Error validating request body")
		return
	}

	err := h.service.SaveItem(ctx, item)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Error saving item")
		return
	}
	h.respondWithJson(w, http.StatusCreated, map[string]interface{}{"item": item})
}

func (h *MealPlanHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data, err := h.service.GetDashboard(ctx)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Error getting dashboard")
		return
	}
	h.respondWithJson(w, http.StatusOK, data)
}
