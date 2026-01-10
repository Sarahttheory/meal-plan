package service

import (
	"meal-plan/internal/models"
)

type Repository interface {
	GetDishes() ([]models.Dish, error)
	SaveDish(dish models.Dish) error
	GetIngredients() ([]models.Ingredient, error)
	GetWeeklyPlan() (models.WeeklyPlan, error)
	SaveWeeklyPlan(plan models.PlanItem) error
}
type MealPlanService struct {
	Repo Repository
}

func NewMealPlanService(repo Repository) *MealPlanService {
	return &MealPlanService{
		Repo: repo,
	}
}
