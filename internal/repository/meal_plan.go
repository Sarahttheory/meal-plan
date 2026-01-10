package repository

import (
	"fmt"
	"meal-plan/internal/models"
)

func (r *MealPlanRepository) GetWeeklyPlan() (models.WeeklyPlan, error) {
	query := `
    SELECT
        d.name as day_name,
        mt.name as meal_type_name,
        di.name as dish_name,
        di.calories
    FROM meal_plan mp
    JOIN days_of_week d ON mp.day_id = d.id
    JOIN meal_types mt ON mp.meal_type_id = mt.id
    JOIN dishes di ON mp.dish_id = di.id
    ORDER BY d.id, mt.sort_order;
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plan := make(models.WeeklyPlan)
	for rows.Next() {
		var dayName string
		var detail models.MealDetail
		err = rows.Scan(&dayName, &detail.MealType, &detail.DishName, &detail.Calories)
		if err != nil {
			return nil, err
		}
		plan[dayName] = append(plan[dayName], detail)
	}
	return plan, nil
}

func (r *MealPlanRepository) SaveWeeklyPlan(plan models.PlanItem) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("repo: SaveWeeklyPlan failed. Could not begin transaction: %w", err)
	}

	_, err = tx.Exec(
		"INSERT INTO meal_plan_weekly_plan (day_id, meal_type_id, dish_id) VALUES ($1, $2, $3)",
		plan.DayId, plan.MealTypeId, plan.DishId,
	)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("repo: SaveWeeklyPlan failed. Could not save meal plan: %w", err)
	}

	return tx.Commit()
}
