package repositories

import (
	"culineira-backend/helpers/auth"
	"culineira-backend/helpers/converter"
	"culineira-backend/helpers/response"
	"culineira-backend/migrations"
	"culineira-backend/modules/users/models"
	"database/sql"
	"fmt"
	"net/http"
)

func GetUserByUsernamePassword(username, password string) (response.LoginResponse, error) {
	var res response.LoginResponse
	var obj models.UserDetail
	var arrobj []models.UserDetail

	var UpdatedAt sql.NullString
	var id string
	var token string
	var msg string = "Failed to get data"

	sql := fmt.Sprintf("SELECT id, country_code, slug_name, username, concat(first_name, last_name) as full_name, email, job, created_at, updated_at "+
		"FROM users WHERE username = '%s' AND password = '%s' LIMIT 1", username, password)

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&id,
			&obj.CountryCode,
			&obj.SlugName,
			&obj.UserName,
			&obj.FullName,
			&obj.Email,
			&obj.Job,
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

	if id != "" {
		token, err = auth.GetToken(id)
		msg = "Successfully get data"

		if err != nil {
			return res, err
		}
	}

	res.Status = http.StatusOK
	res.Message = msg
	res.Data = arrobj
	res.Token = token

	return res, nil
}

func GetUserByUsername(username string) bool {
	sql := fmt.Sprintf("SELECT id "+
		"FROM users WHERE username = '%s' LIMIT 1", username)

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return true
		}
		return true
	}

	return false
}
