package handlers

import (
	"culineira-backend/migrations"
	"culineira-backend/modules/countries/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCountries(c *gin.Context) {
	var (
		result gin.H
	)

	models, err := repositories.GetAllCountries(migrations.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": models,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetAllCountriesRecipe(c *gin.Context) {
	var (
		result gin.H
	)

	models, err := repositories.GetAllCountriesRecipe(migrations.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": models,
		}
	}

	c.JSON(http.StatusOK, result)
}
