package models

type (
	Countries struct {
		CountryCode      string `json:"country_code"`
		CountryName      string `json:"country_name"`
		CountryContinent string `json:"country_continent"`
	}
	CountriesRecipe struct {
		Countries
		RecipeName string `json:"recipe_name"`
		RecipeType string `json:"recipe_type"`
		UserName   string `json:"username"`
	}
)
