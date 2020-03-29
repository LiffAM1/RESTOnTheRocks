package dtos

type Recipe struct {
	Id string
	Name, Url string
	RecipeIngredients []*RecipeIngredient
}

type RecipeIngredient struct {
	Id string
	Ingredient *Ingredient
	Amount float64
	Unit string // We convert all measurement units to ounces
}

type Ingredient struct {
	Id string
	Name string
	Recipes []string // List of Recipe Id's
}

