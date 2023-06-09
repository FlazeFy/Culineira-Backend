package models

type Ingredients struct {
	Id               string `json:"id"`
	RecipeId         string `json:"recipes_id"`
	IngredientName   string `json:"ingredient_name"`
	IngredientVolume string `json:"ingredient_volume"`

	//Properties
	IsOptional bool   `json:"is_optional"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	UpdatedAt  string `json:"updated_at"`
	UpdatedBy  string `json:"updated_by"`
}
