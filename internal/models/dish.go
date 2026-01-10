package models

type Dish struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Calories    int          `json:"calories"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Calories int    `json:"calories"`
}
