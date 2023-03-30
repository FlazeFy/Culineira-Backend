package models

type (
	Ingredients struct {
		IngredientName   string `json:"ingredient_name"`
		IngredientVolume string `json:"ingredient_volume"`

		//Properties
		IsOptional bool   `json:"is_optional"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	}
	CreateIngredient struct {
		Id               string `json:"id"`
		RecipeId         string `json:"recipes_id"`
		IngredientName   string `json:"ingredient_name"`
		IngredientVolume string `json:"ingredient_volume"`

		//Properties
		IsOptional bool   `json:"is_optional"`
		CreatedBy  string `json:"created_by"`
	}
	UpdateIngredient struct {
		IngredientName   string `json:"ingredient_name"`
		IngredientVolume string `json:"ingredient_volume"`

		//Properties
		IsOptional bool   `json:"is_optional"`
		UpdatedBy  string `json:"updated_by"`
	}
)
