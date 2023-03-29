package models

type Recipes struct {
	Id              string `json:"id"`
	CountryCode     string `json:"country_code"`
	RecipeSlug      string `json:"recipe_slug"`
	RecipeName      string `json:"recipe_name"`
	RecipeDesc      string `json:"recipe_desc"`
	RecipeType      string `json:"recipe_type"`
	RecipeTimeSpend int16  `json:"recipe_time_spend"`
	RecipeCal       int16  `json:"recipe_calorie"`
	RecipeLevel     string `json:"recipe_level"`
	RecipeImageURL  string `json:"recipe_image_url"`
	RecipeVideoURL  string `json:"recipe_video_url"`

	//Properties
	IsPrivate bool   `json:"is_private"`
	CreatedAt int64  `json:"created_at"`
	CreatedBy string `json:"created_by"` //user id
	UpdatedAt int64  `json:"updated_at"`
	UpdatedBy string `json:"updated_by"` //user id
	DeletedAt int64  `json:"deleted_at"`
	DeletedBy string `json:"deleted_by"` //user id
}
