package handlers

import (
	"culineira-backend/modules/recipes/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRecipe(c *gin.Context) {
	models, err := repositories.GetAllRecipe()

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, models)
}

func GetRecipeBySlug(c *gin.Context) {
	slug := c.Param("slug")

	models, err := repositories.GetRecipeBySlug(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, models)
}

func GetRecipeDetailBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var (
		result gin.H
	)

	ing, err_c := repositories.GetRecipeIngredientBySlug(slug)
	steps, err_l := repositories.GetRecipeStepBySlug(slug)
	total := "Showing " + strconv.Itoa(ing.Total) + " ingredients and " + strconv.Itoa(steps.Total) + " steps"

	if err_c != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err_c.Error()})
	}
	if err_l != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err_l.Error()})
	}

	result = gin.H{
		"ingredient": ing,
		"step":       steps,
		"total":      total,
	}

	c.JSON(http.StatusOK, result)
}
