package dtos

type Recipe struct {
	Id int
	Name, Url string
	//RecipeIngredients []*RecipeIngredient
}

// RecipeIngredient represents an ingredient that's used in a recipe
type RecipeIngredient struct {
	Id int
	Ingredient *Ingredient
	Amount float32
	Unit string // We convert all measurement units to ounces
}

// IngredientMembership maintains the mapping of ingredients in each recipe
type IngredientMembership struct {
	IngredientId int
	RecipeId int
}

type Ingredient struct {
	Id int
	Name string
	Recipes []string // List of Recipe Id's
}

