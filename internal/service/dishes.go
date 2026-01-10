package service

import (
	"errors"
	"meal-plan/internal/models"
)

func (s *MealPlanService) GetDishes() ([]models.Dish, error) {
	return s.Repo.GetDishes()
}

func (s *MealPlanService) SaveDish(dish models.Dish) error {
	if dish.Name == "" {
		return errors.New("dish name is required")
	}
	return s.Repo.SaveDish(dish)
}

func (s *MealPlanService) GetIngredients() ([]models.Ingredient, error) {
	return s.Repo.GetIngredients()
}
