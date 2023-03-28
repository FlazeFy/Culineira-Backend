package models

type (
	Countries struct {
		Id               string `json:"id"`
		CountryCode      string `json:"country_code"`
		CountryName      string `json:"country_name"`
		CountryContinent string `json:"country_continent"`

		//Properties
		IsActive  bool  `json:"is_active"`
		CreatedAt int64 `json:"created_at"`
		UpdatedAt int64 `json:"updated_at"`
	}
	CountriesRecipe struct {
		CountryCode      string `json:"country_code"`
		CountryName      string `json:"country_name"`
		CountryContinent string `json:"country_continent"`
		RecipeName       string `json:"recipe_name"`
		RecipeType       string `json:"recipe_type"`
		UserName         string `json:"username"`
	}
)
