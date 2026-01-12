package service

import (
	"context"
	"meal-plan/internal/models"
)

func (s *MealPlanService) GetWeeklyPlan(ctx context.Context) (models.WeeklyPlan, error) {
	return s.repo.GetWeeklyPlan(ctx)
}

func (s *MealPlanService) SaveWeeklyPlan(ctx context.Context, plan models.PlanItem) error {
	return s.repo.SaveWeeklyPlan(ctx, plan)
}
