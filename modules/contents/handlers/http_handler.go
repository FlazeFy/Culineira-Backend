package handlers

import (
	"culineira-backend/modules/contents/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetRecipeContent(c *gin.Context) {
	recipe_id := c.Param("recipes_id")
	var (
		result gin.H
	)

	comments, err_c := repositories.GetRecipeComment(recipe_id)
	likes, err_l := repositories.GetRecipeLike(recipe_id)
	total := "Showing " + strconv.Itoa(comments.Total) + " comments and " + strconv.Itoa(likes.Total) + " likes"

	if err_c != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err_c.Error()})
	}
	if err_l != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err_l.Error()})
	}

	result = gin.H{
		"comment": comments,
		"like":    likes,
		"total":   total,
	}

	c.JSON(http.StatusOK, result)
}
