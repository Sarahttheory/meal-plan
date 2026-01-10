package models

type MealDetail struct {
	MealType string `json:"meal_type"`
	DishName string `json:"dish_name"`
	Calories int    `json:"calories"`
}

type WeeklyPlan map[string][]MealDetail
