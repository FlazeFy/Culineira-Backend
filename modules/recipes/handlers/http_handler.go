package handlers

import (
	"culineira-backend/modules/recipes/models"
	"culineira-backend/modules/recipes/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Query
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

// Command
func CreateStep(c *gin.Context) {
	var step models.CreateStep

	err := c.ShouldBindJSON(&step)
	if err != nil {
		panic(err)
	}

	res, err := repositories.CreateStep(step)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}

func CreateIngredient(c *gin.Context) {
	var ing models.CreateIngredient

	err := c.ShouldBindJSON(&ing)
	if err != nil {
		panic(err)
	}

	res, err := repositories.CreateIngredient(ing)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}

func DeleteStepById(c *gin.Context) {
	id := c.Param("id")

	res, err := repositories.DeleteStepById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}

func DeleteIngredientById(c *gin.Context) {
	id := c.Param("id")

	res, err := repositories.DeleteIngredientById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}

func DestroyRecipeById(c *gin.Context) {
	id := c.Param("id")

	res, err := repositories.DestroyRecipeById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}
