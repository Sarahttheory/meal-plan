package models

type Dish struct {
	ID          int          `json:"id"`
	Name        string       `json:"name" validate:"required,min=3,max=100"`
	Calories    int          `json:"calories"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type CreateDishInput struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
	//Calories      int    `json:"calories" validate:"required,min=0"` todo вернуть калории и ввести их подсчет из ингредиентов
	IngredientIds []int `json:"ingredient_ids" validate:"required,min=1"`
}
