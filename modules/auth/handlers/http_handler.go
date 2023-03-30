package handlers

import (
	"culineira-backend/modules/auth/models"
	"culineira-backend/modules/auth/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var login models.UserLogin

	err := c.ShouldBindJSON(&login)
	if err != nil {
		panic(err)
	}

	res, err := repositories.GetUserByUsernamePassword(login.Username, login.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username or Password is incorrect"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func Register(c *gin.Context) {

	var regis models.CreateUser

	if err := c.ShouldBindJSON(&regis); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status := repositories.GetUserByUsername(regis.UserName)

	if !status {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exist"})
		return

	} else {
		res, err := repositories.CreateUser(regis)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}
