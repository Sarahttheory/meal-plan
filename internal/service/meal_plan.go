package service

import (
	"meal-plan/internal/models"
)

func (s *MealPlanService) GetWeeklyPlan() (models.WeeklyPlan, error) {
	return s.repo.GetWeeklyPlan()
}

func (s *MealPlanService) SaveWeeklyPlan(plan models.PlanItem) error {
	return s.repo.SaveWeeklyPlan(plan)
}
