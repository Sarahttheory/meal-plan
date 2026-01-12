package repository

import (
	"context"
	"fmt"
	"meal-plan/internal/models"
)

func (r *MealPlanRepository) GetWeeklyPlan(ctx context.Context) (models.WeeklyPlan, error) {
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
	rows, err := r.DB.QueryContext(ctx, query)
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

func (r *MealPlanRepository) SaveItem(ctx context.Context, item models.PlanItem) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("repo: SaveItem failed. Could not begin transaction: %w", err)
	}
	defer tx.Rollback()

	query := `
        INSERT INTO meal_plan (day_id, meal_type_id, dish_id)
        VALUES ($1, $2, $3)
        ON CONFLICT (day_id, meal_type_id) 
        DO UPDATE SET dish_id = EXCLUDED.dish_id;
    `
	_, err = tx.ExecContext(ctx, query, item.DayId, item.MealTypeId, item.DishId)
	if err != nil {
		return fmt.Errorf("repo: SaveItem failed. Could not insert into MealPlan: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("repo: SaveItem failed. Could not commit transaction: %w", err)
	}

	return nil
}
