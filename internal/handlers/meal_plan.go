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

func (h *MealPlanHandler) SaveWeeklyPlan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var plan models.PlanItem
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Error parsing request body")
		return
	}
	if err := h.validator.Struct(plan); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "Error validating request body")
		return
	}

	err := h.service.SaveWeeklyPlan(ctx, plan)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, "Error saving weekly plan")
		return
	}
	h.respondWithJson(w, http.StatusCreated, map[string]interface{}{"weekly_plan": plan}) //todo проверить
}
