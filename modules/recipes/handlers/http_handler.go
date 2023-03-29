package handlers

import (
	"culineira-backend/modules/recipes/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRecipe(c *gin.Context) {
	models, err := repositories.GetAllRecipe()

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, models)
}

func GetAllRecipeBySlug(c *gin.Context) {
	slug := c.Param("slug")

	models, err := repositories.GetAllRecipeBySlug(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, models)
}
