package repositories

import (
	"culineira-backend/helpers/converter"
	"culineira-backend/helpers/response"
	"culineira-backend/migrations"
	"culineira-backend/modules/users/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
)

func GetAllUser() (response.Response, error) {
	var res response.Response
	var obj models.Users
	var arrobj []models.Users

	sql := "SELECT country_code, slug_name, username, concat(first_name, ' ' ,last_name) as full_name, created_at FROM users WHERE is_active ORDER BY username ASC "

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&obj.CountryCode,
			&obj.SlugName,
			&obj.UserName,
			&obj.FullName,
			&obj.CreatedAt,
		)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(arrobj)) + " data"
	res.Data = arrobj

	return res, nil
}

func GetUserBySlug(slug string) (response.Response, error) {
	var res response.Response
	var obj models.UserDetail
	var arrobj []models.UserDetail

	var UpdatedAt sql.NullString

	sql := fmt.Sprintf("SELECT country_code, slug_name, username, concat(first_name, last_name) as full_name, email, job, created_at, updated_at "+
		"FROM users WHERE is_active AND slug_name = '%s' LIMIT 1", slug)

	rows, err := migrations.DbConnection.Query(sql)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
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

	res.Status = http.StatusOK
	res.Message = "Successfully collect " + strconv.Itoa(len(arrobj)) + " data"
	res.Data = arrobj

	return res, nil
}
