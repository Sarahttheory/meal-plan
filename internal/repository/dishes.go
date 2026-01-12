package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"meal-plan/internal/models"
)

func (r *MealPlanRepository) GetDishes(ctx context.Context) ([]models.Dish, error) {
	query := `
    SELECT d.id, d.name, d.calories,
        COALESCE(json_agg(json_build_object('id', i.id, 'name', i.name)) FILTER (WHERE i.id IS NOT NULL), '[]')
    FROM dishes d
    LEFT JOIN ingredients_in_dish iid ON d.id = iid.dish_id
    LEFT JOIN ingredients i ON iid.ingredient_id = i.id
    GROUP BY d.id;
    `
	rows, err := r.DB.QueryContext(ctx, query)
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
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return dishes, nil
}

func (r *MealPlanRepository) SaveDish(ctx context.Context, dish models.CreateDishInput) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("repo: SaveDish failed. Could not start transaction: %w", err)
	}
	defer tx.Rollback()

	var lastInsertId int
	err = tx.QueryRowContext(ctx,
		"INSERT INTO dishes (name, calories) VALUES ($1, $2) RETURNING id",
		dish.Name, dish.Calories,
	).Scan(&lastInsertId)

	if err != nil {
		return fmt.Errorf("repo: SaveDish failed. Could not insert dish: %w", err)
	}

	for _, ingredientId := range dish.IngredientIds {
		_, err = tx.ExecContext(ctx,
			"INSERT INTO ingredients_in_dish (dish_id, ingredient_id) VALUES ($1, $2)",
			lastInsertId, ingredientId,
		)
		if err != nil {
			return fmt.Errorf("repo: SaveDish failed. Could not insert ingredient: %w", err)
		}
	}
	return tx.Commit()
}

func (r *MealPlanRepository) GetIngredients(ctx context.Context) ([]models.Ingredient, error) {
	query := `
    SELECT id, name
    FROM ingredients;
    `
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []models.Ingredient
	for rows.Next() {
		var ingredient models.Ingredient
		if err := rows.Scan(&ingredient.ID, &ingredient.Name); err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ingredient)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ingredients, nil
}
