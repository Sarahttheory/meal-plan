package handlers

import (
	"meal-plan/internal/models"
	"meal-plan/internal/service"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	GetDishes() ([]models.Dish, error)
	SaveDish(dish models.CreateDishInput) error
	GetIngredients() ([]models.Ingredient, error)
	GetWeeklyPlan() (models.WeeklyPlan, error)
	SaveWeeklyPlan(plan models.PlanItem) error
}

type MealPlanHandler struct {
	service   Service
	validator *validator.Validate
}

func NewMealPlanHandler(srv *service.MealPlanService) *MealPlanHandler {
	return &MealPlanHandler{
		service:   srv,
		validator: validator.New(),
	}
}
