package repository

import "database/sql"

type MealPlanRepository struct {
	DB *sql.DB
}

func NewMealPlanRepository(db *sql.DB) *MealPlanRepository {
	return &MealPlanRepository{
		DB: db,
	}
}
