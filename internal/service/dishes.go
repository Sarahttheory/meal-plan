package service

import (
	"context"
	"meal-plan/internal/models"
)

func (s *MealPlanService) GetDishes(ctx context.Context) ([]models.Dish, error) {
	return s.repo.GetDishes(ctx)
}

func (s *MealPlanService) SaveDish(ctx context.Context, dish models.CreateDishInput) error {
	return s.repo.SaveDish(ctx, dish)
}

func (s *MealPlanService) GetIngredients(ctx context.Context) ([]models.Ingredient, error) {
	return s.repo.GetIngredients(ctx)
}

func (s *MealPlanService) SaveIngredient(ctx context.Context, ingredient models.Ingredient) error {
	return s.repo.SaveIngredient(ctx, ingredient)
}
