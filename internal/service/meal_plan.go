package service

import (
	"meal-plan/internal/models"
)

func (s *MealPlanService) GetWeeklyPlan() (models.WeeklyPlan, error) {
	return s.Repo.GetWeeklyPlan()
}

func (s *MealPlanService) SaveWeeklyPlan(plan models.PlanItem) error {
	//TODO валидацию докинуть
	return s.Repo.SaveWeeklyPlan(plan)
}
