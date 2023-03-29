package repositories

import (
	"culineira-backend/helpers/converter"
	"culineira-backend/helpers/response"
	"culineira-backend/migrations"
	"culineira-backend/modules/recipes/models"
	"database/sql"
	"fmt"
	"net/http"
)

func GetRecipeStepBySlug(slug string) (response.ContentResponse, error) {
	var res response.ContentResponse
	var obj models.Steps
	var arrobj []models.Steps

	var StepsImageUrl sql.NullString
	var StepsVideoUrl sql.NullString
	var UpdatedAt sql.NullString

	sql := fmt.Sprintf("SELECT sort_number, steps_body, steps_image_url, steps_video_url, is_optional, s.created_at, s.updated_at "+
		"FROM steps s JOIN recipes r ON s.recipes_id = r.id "+
		"WHERE r.recipe_slug = '%s'  ORDER BY sort_number ASC ", slug)

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.SortNumber,
			&obj.StepsBody,
			&StepsImageUrl,
			&StepsVideoUrl,
			&obj.IsOptional,
			&obj.CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return res, err
		}

		//Null validator
		obj.StepsImageUrl = converter.CheckNullString(StepsImageUrl)
		obj.StepsVideoUrl = converter.CheckNullString(StepsVideoUrl)
		obj.UpdatedAt = converter.CheckNullString(UpdatedAt)

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect data"
	res.Data = arrobj
	res.Total = len(arrobj)

	return res, nil
}
