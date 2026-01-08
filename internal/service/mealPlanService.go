package service

import (
	"meal-plan/internal/models"
	"meal-plan/internal/repository"
)

type MealPlanService struct {
	Repo *repository.MealPlanRepository
}

func (s *MealPlanService) GetDishes() ([]models.Dish, error) {
	return s.Repo.GetAllDishes()
}
