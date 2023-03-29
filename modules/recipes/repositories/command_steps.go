package repositories

import (
	"culineira-backend/helpers/response"
	"culineira-backend/migrations"
	"fmt"
	"net/http"
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
