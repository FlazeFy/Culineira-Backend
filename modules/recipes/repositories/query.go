package repositories

import (
	"culineira-backend/helpers/converter"
	"culineira-backend/helpers/response"
	"culineira-backend/migrations"
	"culineira-backend/modules/recipes/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func GetAllRecipe() (response.Response, error) {
	var res response.Response
	var obj models.Recipes
	var arrobj []models.Recipes
	var RecipeImageURL sql.NullString
	var RecipeVideoURL sql.NullString
	var UpdatedAt sql.NullString
	var UpdatedBy sql.NullString

	sql := "SELECT country_code, recipe_slug, recipe_name, recipe_desc, recipe_type, recipe_time_spend, recipe_cal, recipe_level, recipe_image_url, recipe_video_url, is_private, created_at, created_by, updated_at, updated_by " +
		"FROM recipes WHERE deleted_at is null and deleted_by is null ORDER BY created_at ASC "

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.CountryCode,
			&obj.RecipeSlug,
			&obj.RecipeName,
			&obj.RecipeDesc,
			&obj.RecipeType,
			&obj.RecipeTimeSpend,
			&obj.RecipeCal,
			&obj.RecipeLevel,
			&RecipeImageURL,
			&RecipeVideoURL,
			&obj.IsPrivate,
			&obj.CreatedAt,
			&obj.CreatedBy,
			&UpdatedAt,
			&UpdatedBy,
		)
		if err != nil {
			return res, err
		}

		//Null validator
		obj.RecipeImageURL = converter.CheckNullString(RecipeImageURL)
		obj.RecipeVideoURL = converter.CheckNullString(RecipeVideoURL)
		obj.UpdatedAt = converter.CheckNullString(UpdatedAt)
		obj.UpdatedBy = converter.CheckNullString(UpdatedBy)

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(arrobj)) + " data"
	res.Data = arrobj

	return res, nil
}

func GetRecipeBySlug(slug string) (response.Response, error) {
	var res response.Response
	var obj models.Recipes
	var arrobj []models.Recipes

	var RecipeImageURL sql.NullString
	var RecipeVideoURL sql.NullString
	var UpdatedAt sql.NullString
	var UpdatedBy sql.NullString

	sql := fmt.Sprintf("SELECT country_code, recipe_slug, recipe_name, recipe_desc, recipe_type, recipe_time_spend, recipe_cal, recipe_level, recipe_image_url, recipe_video_url, is_private, created_at, created_by, updated_at, updated_by "+
		"FROM recipes WHERE deleted_at is null and deleted_by is null and recipe_slug = '%s'", slug)

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.CountryCode,
			&obj.RecipeSlug,
			&obj.RecipeName,
			&obj.RecipeDesc,
			&obj.RecipeType,
			&obj.RecipeTimeSpend,
			&obj.RecipeCal,
			&obj.RecipeLevel,
			&RecipeImageURL,
			&RecipeVideoURL,
			&obj.IsPrivate,
			&obj.CreatedAt,
			&obj.CreatedBy,
			&UpdatedAt,
			&UpdatedBy,
		)
		if err != nil {
			return res, err
		}

		//Null validator
		obj.RecipeImageURL = converter.CheckNullString(RecipeImageURL)
		obj.RecipeVideoURL = converter.CheckNullString(RecipeVideoURL)
		obj.UpdatedAt = converter.CheckNullString(UpdatedAt)
		obj.UpdatedBy = converter.CheckNullString(UpdatedBy)

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(arrobj)) + " data"
	res.Data = arrobj

	return res, nil
}
