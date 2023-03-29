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

func GetRecipeIngredientBySlug(slug string) (response.ContentResponse, error) {
	var res response.ContentResponse
	var obj models.Ingredients
	var arrobj []models.Ingredients

	var UpdatedAt sql.NullString

	sql := fmt.Sprintf("SELECT ingredient_name, ingredient_volume, is_optional, i.created_at, i.updated_at "+
		"FROM ingredients i JOIN recipes r ON i.recipes_id = r.id "+
		"WHERE r.recipe_slug = '%s'", slug)

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&obj.IngredientName,
			&obj.IngredientVolume,
			&obj.IsOptional,
			&obj.CreatedAt,
			&UpdatedAt,
		)
		if err != nil {
			return res, err
		}

		//Null validator
		obj.UpdatedAt = converter.CheckNullString(UpdatedAt)

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect data"
	res.Data = arrobj
	res.Total = len(arrobj)

	return res, nil
}
