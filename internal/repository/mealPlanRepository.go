package repository

import (
  "database/sql"
  "meal-plan/internal/models"
)

type MealPlanRepository struct {
  DB *sql.DB
}

func (r *MealPlanRepository) GetAllDishes() ([]models.Dish, error) {
  query := `SELECT id, name, calories FROM dishes`
  rows, err := r.DB.Query(query)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var dishes []models.Dish
  for rows.Next() {
    var dish models.Dish
    if err := rows.Scan(&dish.ID, &dish.Name, &dish.Calories); err != nil {
      return nil, err
    }
    dishes = append(dishes, dish)
  }
  return dishes, nil
}
