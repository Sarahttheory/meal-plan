package service

import (
	"context"
	"meal-plan/internal/models"
)

type Repository interface {
	GetDishes(ctx context.Context) ([]models.Dish, error)
	SaveDish(ctx context.Context, dish models.CreateDishInput) error
	GetIngredients(ctx context.Context) ([]models.Ingredient, error)
	GetWeeklyPlan(ctx context.Context) (models.WeeklyPlan, error)
	SaveItem(ctx context.Context, item models.PlanItem) error
	SaveIngredient(ctx context.Context, ingredient models.Ingredient) error
}
type MealPlanService struct {
	repo Repository
}

func NewMealPlanService(repo Repository) *MealPlanService {
	return &MealPlanService{
		repo: repo,
	}
}
