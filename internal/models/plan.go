package models

type MealDetail struct {
	MealType string `json:"meal_type"`
	DishName string `json:"dish_name"`
	Calories int    `json:"calories"`
}

type WeeklyPlan map[string][]MealDetail

type PlanItem struct {
	ID         int `json:"id"`
	DayId      int `json:"day_id" validate:"required,gt=0"`
	MealTypeId int `json:"meal_type_id" validate:"required,gt=0"`
	DishId     int `json:"dish_id" validate:"required,gt=0"`
}

type DashboardData struct {
	Dishes      []Dish       `json:"dishes"`
	Ingredients []Ingredient `json:"ingredients"`
	WeeklyPlan  WeeklyPlan   `json:"weekly_plan"`
}
