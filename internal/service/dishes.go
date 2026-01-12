package service

import (
	"meal-plan/internal/models"
)

func (s *MealPlanService) GetDishes() ([]models.Dish, error) {
	return s.repo.GetDishes()
}

func (s *MealPlanService) SaveDish(dish models.CreateDishInput) error {
	return s.repo.SaveDish(dish)
}

func (s *MealPlanService) GetIngredients() ([]models.Ingredient, error) {
	return s.repo.GetIngredients()
}
