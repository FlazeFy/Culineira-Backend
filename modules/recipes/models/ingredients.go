package models

type Ingredients struct {
	IngredientName   string `json:"ingredient_name"`
	IngredientVolume string `json:"ingredient_volume"`

	//Properties
	IsOptional bool   `json:"is_optional"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
