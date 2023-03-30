package repositories

import (
	"culineira-backend/helpers/generator"
	"culineira-backend/helpers/response"
	"culineira-backend/migrations"
	"culineira-backend/modules/contents/models"
	"fmt"
	"net/http"
	"strings"
)

func CreateLike(step models.CreateLikes) (response.Response, error) {
	var res response.Response

	createdAt := generator.GetDateTimeNow()
	id := generator.GetUUID()
	recipeId := strings.ReplaceAll(step.RecipeId, "'", "''")
	createdBy := strings.ReplaceAll(step.CreatedBy, "'", "''")

	sql := fmt.Sprintf("INSERT INTO likes "+
		"(id, recipes_id, created_at, created_by) "+
		"VALUES('%s', '%s', '%s', '%s') ",
		id, recipeId, createdAt, createdBy,
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

func CreateComment(step models.CreateComments) (response.Response, error) {
	var res response.Response

	createdAt := generator.GetDateTimeNow()
	id := generator.GetUUID()
	recipeId := strings.ReplaceAll(step.RecipeId, "'", "''")
	cmntBody := strings.ReplaceAll(step.CommentBody, "'", "''")
	cmntFileUrl := strings.ReplaceAll(step.CommentFileUrl, "'", "''")
	createdBy := strings.ReplaceAll(step.CreatedBy, "'", "''")

	sql := fmt.Sprintf("INSERT INTO comments "+
		"(id, recipes_id, comment_body, comment_file_url, created_at, created_by, deleted_at, deleted_by) "+
		"VALUES('%s', '%s', '%s', '%s', '%s', '%s', null, null) ",
		id, recipeId, cmntBody, cmntFileUrl, createdAt, createdBy,
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

func DestroyCommentById(id string) (response.Response, error) {
	var res response.Response

	sql := fmt.Sprintf("DELETE FROM comments WHERE id = '%s'", id)

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

func DestroyLikeById(id string) (response.Response, error) {
	var res response.Response

	sql := fmt.Sprintf("DELETE FROM likes WHERE id = '%s'", id)

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
