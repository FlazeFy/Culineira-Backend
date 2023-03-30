package handlers

import (
	"culineira-backend/modules/contents/models"
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

func CreateLike(c *gin.Context) {
	var like models.CreateLikes

	err := c.ShouldBindJSON(&like)
	if err != nil {
		panic(err)
	}

	res, err := repositories.CreateLike(like)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}

func CreateComment(c *gin.Context) {
	var cmnt models.CreateComments

	err := c.ShouldBindJSON(&cmnt)
	if err != nil {
		panic(err)
	}

	res, err := repositories.CreateComment(cmnt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, res)
}
