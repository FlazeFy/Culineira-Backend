package repositories

import (
	"culineira-backend/helpers/generator"
	"culineira-backend/helpers/response"
	"culineira-backend/migrations"
	"culineira-backend/modules/recipes/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func DestroyRecipeById(id string) (response.Response, error) {
	var res response.Response

	sql := fmt.Sprintf("DELETE FROM recipes WHERE id = '%s'", id)

	result, err := migrations.DbConnection.Exec(sql)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = fmt.Sprintf("Successfully removed %d data", rowsAffected)
	res.Data = nil

	return res, nil
}

func CreateRecipe(step models.CreateRecipe) (response.Response, error) {
	var res response.Response

	createdAt := generator.GetDateTimeNow()
	id := generator.GetUUID()
	slug := generator.GetSlug(step.RecipeName)

	countryCode := strings.ReplaceAll(step.CountryCode, "'", "''")
	recipeName := strings.ReplaceAll(step.RecipeName, "'", "''")
	recipeDesc := strings.ReplaceAll(step.RecipeDesc, "'", "''")
	recipeType := strings.ReplaceAll(step.RecipeType, "'", "''")
	recipeLevel := strings.ReplaceAll(step.RecipeLevel, "'", "''")
	recipeImage := strings.ReplaceAll(step.RecipeImageURL, "'", "''")
	recipeVideo := strings.ReplaceAll(step.RecipeVideoURL, "'", "''")
	createdBy := strings.ReplaceAll(step.CreatedBy, "'", "''")
	IsPrivate := strconv.FormatBool(step.IsPrivate)

	sql := fmt.Sprintf("INSERT INTO recipes "+
		"(id, country_code, recipe_name, recipe_desc, recipe_type, recipe_time_spend, recipe_cal, recipe_level, recipe_image_url, recipe_video_url, is_private, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, recipe_slug) "+
		"VALUES('%s', '%s', '%s', '%s', '%s', %d, %d, '%s', '%s', '%s', '%s', '%s', '%s', NULL, NULL, NULL, NULL, '%s') ",
		id, countryCode, recipeName, recipeDesc, recipeType, step.RecipeTimeSpend, step.RecipeCal, recipeLevel, recipeImage, recipeVideo, IsPrivate, createdAt, createdBy, slug,
	)

	result, err := migrations.DbConnection.Exec(sql)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = fmt.Sprintf("Successfully inserted %d data", rowsAffected)
	res.Data = fmt.Sprintf("Your data ID : %s", id)

	return res, nil
}

func UpdateRecipeById(step models.UpdateRecipe, id string) (response.Response, error) {
	var res response.Response

	updatedAt := generator.GetDateTimeNow()

	countryCode := strings.ReplaceAll(step.CountryCode, "'", "''")
	recipeDesc := strings.ReplaceAll(step.RecipeDesc, "'", "''")
	recipeType := strings.ReplaceAll(step.RecipeType, "'", "''")
	recipeLevel := strings.ReplaceAll(step.RecipeLevel, "'", "''")
	recipeImage := strings.ReplaceAll(step.RecipeImageURL, "'", "''")
	recipeVideo := strings.ReplaceAll(step.RecipeVideoURL, "'", "''")
	updatedBy := strings.ReplaceAll(step.UpdatedBy, "'", "''")
	IsPrivate := strconv.FormatBool(step.IsPrivate)

	sql := fmt.Sprintf("UPDATE recipes "+
		"SET country_code='%s', recipe_desc='%s', recipe_type='%s', recipe_time_spend=%d, recipe_cal=%d, recipe_level='%s', recipe_image_url='%s', recipe_video_url='%s', is_private=%s, updated_at='%s', updated_by='%s' "+
		"WHERE id='%s' ",
		countryCode, recipeDesc, recipeType, step.RecipeTimeSpend, step.RecipeCal, recipeLevel, recipeImage, recipeVideo, IsPrivate, updatedAt, updatedBy, id,
	)

	result, err := migrations.DbConnection.Exec(sql)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = fmt.Sprintf("Successfully updated %d data", rowsAffected)
	res.Data = fmt.Sprintf("Your data ID : %s", id)

	return res, nil
}
