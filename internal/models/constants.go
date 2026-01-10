package models

const (
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
	Sunday    = 7
)

type PlanItem struct {
	ID         int `json:"id"`
	DayId      int `json:"day_id" validate:"required,gt=0"`
	MealTypeId int `json:"meal_type_id" validate:"required,gt=0"`
	DishId     int `json:"dish_id" validate:"required,gt=0"`
}
