package handlers

import (
	"culineira-backend/modules/countries/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCountries(c *gin.Context) {
	models, err := repositories.GetAllCountries()

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, models)
}

func GetAllCountriesRecipe(c *gin.Context) {
	code := c.Param("countrycode")

	models, err := repositories.GetAllCountriesRecipe(code)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, models)
}
