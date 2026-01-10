package handlers

import (
	"meal-plan/internal/models"
	"meal-plan/internal/service"
)

type Service interface {
	GetDishes() ([]models.Dish, error)
	SaveDish(dish models.Dish) error
	GetIngredients() ([]models.Ingredient, error)
	GetWeeklyPlan() (models.WeeklyPlan, error)
	SaveWeeklyPlan(plan models.PlanItem) error
}

type MealPlanHandler struct {
	Service Service
}

func NewMealPlanHandler(srv *service.MealPlanService) *MealPlanHandler {
	return &MealPlanHandler{
		Service: srv,
	}
}
