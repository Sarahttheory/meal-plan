package repository

import (
	"encoding/json"
	"fmt"
	"meal-plan/internal/models"
)

func (r *MealPlanRepository) GetDishes() ([]models.Dish, error) {
	query := `
    SELECT d.id, d.name, d.calories,
        COALESCE(json_agg(json_build_object('id', i.id, 'name', i.name)) FILTER (WHERE i.id IS NOT NULL), '[]')
    FROM dishes d
    LEFT JOIN ingredients_in_dish iid ON d.id = iid.dish_id
    LEFT JOIN ingredients i ON iid.ingredient_id = i.id
    GROUP BY d.id;
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dishes []models.Dish

	for rows.Next() {
		var dish models.Dish
		var ingredients []byte

		if err := rows.Scan(&dish.ID, &dish.Name, &dish.Calories, &ingredients); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(ingredients, &dish.Ingredients); err != nil {
			return nil, err
		}

		dishes = append(dishes, dish)
	}
	return dishes, nil
}

func (r *MealPlanRepository) SaveDish(dish models.Dish) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return fmt.Errorf("repo: SaveDish failed. Could not start transaction: %w", err)
	}

	var lastInsertId int
	err = tx.QueryRow(
		"INSERT INTO dishes (name, calories) VALUES ($1, $2) RETURNING id",
		dish.Name, dish.Calories,
	).Scan(&lastInsertId)

	if err != nil {
		tx.Rollback()
		return fmt.Errorf("repo: SaveDish failed. Could not insert dish: %w", err)
	}

	for _, ingredientId := range dish.Ingredients {
		_, err = tx.Exec(
			"INSERT INTO ingredients_in_dish (dish_id, ingredient_id) VALUES ($1, $2)",
			lastInsertId, ingredientId,
		)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("repo: SaveDish failed. Could not insert ingredient: %w", err)
		}
	}
	return tx.Commit()
}

func (r *MealPlanRepository) GetIngredients() ([]models.Ingredient, error) {
	query := `
    SELECT id, name, calories
    FROM ingredients;
    `
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []models.Ingredient
	for rows.Next() {
		var ingredient models.Ingredient
		if err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Calories); err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}
