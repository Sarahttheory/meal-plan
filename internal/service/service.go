package service

import (
	"meal-plan/internal/models"
)

type Repository interface {
	GetDishes() ([]models.Dish, error)
	SaveDish(dish models.CreateDishInput) error
	GetIngredients() ([]models.Ingredient, error)
	GetWeeklyPlan() (models.WeeklyPlan, error)
	SaveWeeklyPlan(plan models.PlanItem) error
}
type MealPlanService struct {
	repo Repository
}

func NewMealPlanService(repo Repository) *MealPlanService {
	return &MealPlanService{
		repo: repo,
	}
}
