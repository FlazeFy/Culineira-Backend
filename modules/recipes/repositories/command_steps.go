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

func DeleteStepById(id string) (response.Response, error) {
	var res response.Response

	sql := fmt.Sprintf("DELETE FROM steps WHERE id = '%s'", id)

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

func CreateStep(step models.CreateStep) (response.Response, error) {
	var res response.Response

	createdAt := generator.GetDateTimeNow()
	id := generator.GetUUID()
	recipeId := strings.ReplaceAll(step.RecipeId, "'", "''")
	stepsBody := strings.ReplaceAll(step.StepsBody, "'", "''")
	stepsImgUrl := strings.ReplaceAll(step.StepsImageUrl, "'", "''")
	stepsVideoUrl := strings.ReplaceAll(step.StepsVideoUrl, "'", "''")
	CreatedBy := strings.ReplaceAll(step.CreatedBy, "'", "''")
	isOptional := strconv.FormatBool(step.IsOptional)

	sql := fmt.Sprintf("INSERT INTO steps "+
		"(id, recipes_id, steps_body, steps_image_url, steps_video_url, is_optional, created_at, created_by, updated_at, updated_by, sort_number) "+
		"VALUES('%s', '%s', '%s', '%s', '%s', %s, '%s', '%s', NULL, NULL, %d) ",
		id, recipeId, stepsBody, stepsImgUrl, stepsVideoUrl, isOptional, createdAt, CreatedBy, step.SortNumber,
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

func UpdateStep(step models.UpdateStep, id string) (response.Response, error) {
	var res response.Response

	updatedAt := generator.GetDateTimeNow()
	stepBody := strings.ReplaceAll(step.StepsBody, "'", "''")
	stepImage := strings.ReplaceAll(step.StepsImageUrl, "'", "''")
	stepVideo := strings.ReplaceAll(step.StepsVideoUrl, "'", "''")
	updatedBy := strings.ReplaceAll(step.UpdatedBy, "'", "''")
	isOptional := strconv.FormatBool(step.IsOptional)

	sql := fmt.Sprintf("UPDATE steps "+
		"SET steps_body='%s', steps_image_url='%s', steps_video_url='%s', is_optional=%s, updated_at='%s', updated_by='%s' "+
		"WHERE id='%s' ",
		stepBody, stepImage, stepVideo, isOptional, updatedAt, updatedBy, id,
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
