package service

import (
	"context"
	"meal-plan/internal/models"
	"sync"
)

func (s *MealPlanService) GetWeeklyPlan(ctx context.Context) (models.WeeklyPlan, error) {
	return s.repo.GetWeeklyPlan(ctx)
}

func (s *MealPlanService) SaveWeeklyPlan(ctx context.Context, plan models.PlanItem) error {
	return s.repo.SaveWeeklyPlan(ctx, plan)
}

// todo можно будет улучшить с помощью errgroup
func (s *MealPlanService) GetDashboard(ctx context.Context) (models.DashboardData, error) {
	var (
		data    models.DashboardData
		wg      sync.WaitGroup
		errChar = make(chan error, 3)
	)
	wg.Add(3)

	go func() {
		defer wg.Done()
		var err error
		data.Dishes, err = s.repo.GetDishes(ctx)
		if err != nil {
			errChar <- err
		}
	}()
	go func() {
		defer wg.Done()
		var err error
		data.Ingredients, err = s.repo.GetIngredients(ctx)
		if err != nil {
			errChar <- err
		}
	}()
	go func() {
		defer wg.Done()
		var err error
		data.WeeklyPlan, err = s.repo.GetWeeklyPlan(ctx)
		if err != nil {
			errChar <- err
		}
	}()
	wg.Wait()
	close(errChar)

	for err := range errChar {
		if err != nil {
			return data, err
		}
	}
	return data, nil
}
