package repositories

import (
	"culineira-backend/helpers/converter"
	"culineira-backend/helpers/response"
	"culineira-backend/migrations"
	"culineira-backend/modules/contents/models"
	"database/sql"
	"fmt"
	"net/http"
)

func GetRecipeComment(recipe_id string) (response.ContentResponse, error) {
	var res response.ContentResponse
	var obj models.Comments
	var arrobj []models.Comments

	var CommentFileUrl sql.NullString

	sql := fmt.Sprintf("SELECT comment_body, comment_file_url, c.created_at, u.username FROM comments c "+
		"JOIN users u ON u.id = c.created_by "+
		"JOIN recipes r ON r.id = c.recipes_id "+
		"WHERE c.deleted_at is null AND c.deleted_by is null AND recipes_id = '%s' ORDER BY c.created_at DESC", recipe_id)

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&obj.CommentBody,
			&CommentFileUrl,
			&obj.CreatedAt,
			&obj.Username,
		)
		if err != nil {
			return res, err
		}

		//Null validator
		obj.CommentFileUrl = converter.CheckNullString(CommentFileUrl)

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect data"
	res.Data = arrobj
	res.Total = len(arrobj)

	return res, nil
}

func GetRecipeLike(recipe_id string) (response.ContentResponse, error) {
	var res response.ContentResponse
	var obj models.Likes
	var arrobj []models.Likes

	sql := fmt.Sprintf("SELECT l.created_at, u.username FROM likes l "+
		"JOIN users u ON u.id = l.created_by "+
		"JOIN recipes r ON r.id = l.recipes_id "+
		"WHERE recipes_id = '%s' ORDER BY l.created_at DESC", recipe_id)

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&obj.CreatedAt,
			&obj.Username,
		)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect data"
	res.Data = arrobj
	res.Total = len(arrobj)

	return res, nil
}
