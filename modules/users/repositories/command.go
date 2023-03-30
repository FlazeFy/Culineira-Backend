package repositories

import (
	"culineira-backend/helpers/generator"
	"culineira-backend/helpers/response"
	"culineira-backend/migrations"
	"culineira-backend/modules/users/models"
	"fmt"
	"net/http"
	"strings"
)

func DeleteUserById(id string) (response.Response, error) {
	var res response.Response

	sql := fmt.Sprintf("DELETE FROM users WHERE id = '%s'", id)

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

func UpdateUserById(step models.UpdateUser, id string) (response.Response, error) {
	var res response.Response

	updatedAt := generator.GetDateTimeNow()
	countryCode := strings.ReplaceAll(step.CountryCode, "'", "''")
	firstName := strings.ReplaceAll(step.FirstName, "'", "''")
	lastName := strings.ReplaceAll(step.LastName, "'", "''")
	job := strings.ReplaceAll(step.Job, "'", "''")
	updatedBy := strings.ReplaceAll(step.UpdatedBy, "'", "''")

	sql := fmt.Sprintf("UPDATE users "+
		"SET country_code='%s', first_name='%s', last_name='%s', job='%s', updated_at='%s', updated_by='%s' "+
		"WHERE id='%s' ",
		countryCode, firstName, lastName, job, updatedAt, updatedBy, id,
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
