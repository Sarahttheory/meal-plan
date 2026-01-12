package handlers

import (
	"context"
	"meal-plan/internal/models"
	"meal-plan/internal/service"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	GetDishes(ctx context.Context) ([]models.Dish, error)
	SaveDish(ctx context.Context, dish models.CreateDishInput) error
	GetIngredients(ctx context.Context) ([]models.Ingredient, error)
	GetWeeklyPlan(ctx context.Context) (models.WeeklyPlan, error)
	SaveWeeklyPlan(ctx context.Context, plan models.PlanItem) error
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
