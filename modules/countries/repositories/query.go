package repositories

import (
	"culineira-backend/helpers"
	"culineira-backend/migrations"
	"culineira-backend/modules/countries/models"
	"fmt"
	"net/http"
	"strconv"
)

func GetAllCountries() (helpers.Response, error) {
	var res helpers.Response
	var obj models.Countries
	var arrobj []models.Countries

	sql := "SELECT country_code, country_name, country_continent FROM countries WHERE is_active ORDER BY country_name "

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&obj.CountryCode,
			&obj.CountryName,
			&obj.CountryContinent,
		)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(arrobj)) + " data"
	res.Data = arrobj

	return res, nil
}

func GetAllCountriesRecipe(code string) (helpers.Response, error) {
	var res helpers.Response
	var obj models.CountriesRecipe
	var arrobj []models.CountriesRecipe

	sql := fmt.Sprintf("SELECT c.country_code, country_name, country_continent, recipe_name, recipe_type, username as posted_by FROM countries c "+
		"JOIN recipes r on c.country_code = r.country_code JOIN users u on u.id = r.created_by WHERE r.country_code = '%s' "+
		" order by c.country_name ASC", code)

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.CountryCode,
			&obj.CountryName,
			&obj.CountryContinent,
			&obj.RecipeName,
			&obj.RecipeType,
			&obj.UserName,
		)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(arrobj)) + " data"
	res.Data = arrobj

	return res, nil
}
