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

func DeleteIngredientById(id string) (response.Response, error) {
	var res response.Response

	sql := fmt.Sprintf("DELETE FROM ingredients WHERE id = '%s'", id)

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

func CreateIngredient(step models.CreateIngredient) (response.Response, error) {
	var res response.Response

	createdAt := generator.GetDateTimeNow()
	id := generator.GetUUID()
	recipeId := strings.ReplaceAll(step.RecipeId, "'", "''")
	ingName := strings.ReplaceAll(step.IngredientName, "'", "''")
	ingVol := strings.ReplaceAll(step.IngredientVolume, "'", "''")
	CreatedBy := strings.ReplaceAll(step.CreatedBy, "'", "''")
	isOptional := strconv.FormatBool(step.IsOptional)

	sql := fmt.Sprintf("INSERT INTO ingredients "+
		"(id, recipes_id, ingredient_name, ingredient_volume, is_optional, created_at, created_by, updated_at, updated_by) "+
		"VALUES('%s', '%s', '%s', '%s', %s, '%s', '%s', NULL, NULL) ",
		id, recipeId, ingName, ingVol, isOptional, createdAt, CreatedBy,
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
	res.Message = fmt.Sprintf("Successfully insert %d data", rowsAffected)
	res.Data = fmt.Sprintf("Your data ID : %s", id)

	return res, nil
}

func UpdateIngredient(step models.UpdateIngredient, id string) (response.Response, error) {
	var res response.Response

	updatedAt := generator.GetDateTimeNow()
	ingName := strings.ReplaceAll(step.IngredientName, "'", "''")
	ingVol := strings.ReplaceAll(step.IngredientVolume, "'", "''")
	updatedBy := strings.ReplaceAll(step.UpdatedBy, "'", "''")
	isOptional := strconv.FormatBool(step.IsOptional)

	sql := fmt.Sprintf("UPDATE ingredients "+
		"SET ingredient_name='%s', ingredient_volume='%s', is_optional=%s, updated_at='%s', updated_by='%s' "+
		"WHERE id='%s' ",
		ingName, ingVol, isOptional, updatedAt, updatedBy, id,
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
