package models

type (
	RecipeInfo struct {
		RecipeDesc      string `json:"recipe_desc"`
		RecipeType      string `json:"recipe_type"`
		RecipeTimeSpend int16  `json:"recipe_time_spend"`
		RecipeCal       int16  `json:"recipe_calorie"`
		RecipeLevel     string `json:"recipe_level"`
		RecipeImageURL  string `json:"recipe_image_url"`
		RecipeVideoURL  string `json:"recipe_video_url"`
	}
	Recipes struct {
		CountryCode string `json:"country_code"`
		RecipeSlug  string `json:"recipe_slug"`
		RecipeName  string `json:"recipe_name"`
		RecipeInfo

		//Properties
		IsPrivate bool   `json:"is_private"`
		CreatedAt string `json:"created_at"`
		CreatedBy string `json:"created_by"` //user id
		UpdatedAt string `json:"updated_at"`
		UpdatedBy string `json:"updated_by"` //user id
	}
	CreateRecipe struct {
		CountryCode string `json:"country_code"`
		RecipeName  string `json:"recipe_name"`
		RecipeInfo

		//Properties
		IsPrivate bool   `json:"is_private"`
		CreatedBy string `json:"created_by"` //user id
	}
	UpdateRecipe struct {
		CountryCode string `json:"country_code"`
		RecipeInfo

		//Properties
		IsPrivate bool   `json:"is_private"`
		UpdatedBy string `json:"updated_by"` //user id
	}
)
