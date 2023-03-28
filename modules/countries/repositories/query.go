package repositories

import (
	"culineira-backend/modules/countries/models"
	"database/sql"
)

func GetAllCountries(db *sql.DB) (results []models.Countries, err error) {

	sql := "SELECT id, country_code, country_name, country_continent, is_active, created_at, updated_at FROM countries"

	rows, err := db.Query(sql)

	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		var data = models.Countries{}

		err := rows.Scan(
			&data.Id,
			&data.CountryCode,
			&data.CountryName,
			&data.CountryContinent,
			&data.IsActive,
			&data.CreatedAt,
			&data.UpdatedAt,
		)
		if err != nil {
			return results, err
		}
		results = append(results, data)
	}

	return results, nil
}

func GetAllCountriesRecipe(db *sql.DB) (results []models.CountriesRecipe, err error) {

	sql := "SELECT c.country_code, country_name, country_continent, recipe_name, recipe_type, username FROM countries c " +
		"JOIN recipes r on c.country_code = r.country_code JOIN users u on u.id = r.created_by order by c.country_name ASC"

	rows, err := db.Query(sql)

	if err != nil {
		return results, err
	}
	defer rows.Close()

	for rows.Next() {
		var data = models.CountriesRecipe{}

		err := rows.Scan(
			&data.CountryCode,
			&data.CountryName,
			&data.CountryContinent,
			&data.RecipeName,
			&data.RecipeType,
			&data.UserName,
		)
		if err != nil {
			return results, err
		}
		results = append(results, data)
	}

	return results, nil
}
