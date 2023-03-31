package repositories

import (
	"culineira-backend/helpers/auth"
	"culineira-backend/helpers/generator"
	"culineira-backend/helpers/response"
	"culineira-backend/migrations"
	"culineira-backend/modules/auth/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func CreateUser(step models.CreateUser) (response.LoginResponse, error) {
	var res response.LoginResponse
	var token string

	createdAt := generator.GetDateTimeNow()
	id := generator.GetUUID()
	slug := generator.GetSlug(id)

	countryCode := strings.ReplaceAll(step.CountryCode, "'", "''")
	username := strings.ReplaceAll(step.UserName, "'", "''")
	firstName := strings.ReplaceAll(step.FirstName, "'", "''")
	lastName := strings.ReplaceAll(step.LastName, "'", "''")
	job := strings.ReplaceAll(step.Job, "'", "''")
	email := strings.ReplaceAll(step.Email, "'", "''")
	password := strings.ReplaceAll(step.Password, "'", "''")
	isActive := strconv.FormatBool(step.IsActive)

	sql := fmt.Sprintf("INSERT INTO users "+
		"(id, country_code, slug_name, username, first_name, last_name, job, email, password, is_active, created_at, created_by, updated_at, updated_by) "+
		"VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %s, '%s', '%s', NULL, NULL) ",
		id, countryCode, slug, username, firstName, lastName, job, email, password, isActive, createdAt, id,
	)

	result, err := migrations.DbConnection.Exec(sql)

	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	if id != "" {
		token, err = auth.GetToken(id)

		if err != nil {
			return res, err
		}
	}

	res.Status = http.StatusOK
	res.Message = fmt.Sprintf("Successfully insert %d data", rowsAffected)
	res.Data = fmt.Sprintf("Your data ID : %s", id)
	res.Token = token

	return res, nil
}
