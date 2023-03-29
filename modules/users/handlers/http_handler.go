package handlers

import (
	"culineira-backend/modules/users/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	models, err := repositories.GetAllUser()

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, models)
}

func GetUserBySlug(c *gin.Context) {
	slug := c.Param("slug")

	models, err := repositories.GetUserBySlug(slug)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	c.JSON(http.StatusOK, models)
}
