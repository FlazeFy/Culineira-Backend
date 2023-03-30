package main

import (
	"culineira-backend/migrations"
	content_handler "culineira-backend/modules/contents/handlers"
	country_handler "culineira-backend/modules/countries/handlers"
	recipe_handler "culineira-backend/modules/recipes/handlers"
	user_handler "culineira-backend/modules/users/handlers"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "nopass123"
	dbname   = "culineira"
)

func main() {

	psqInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err := sql.Open("postgres", psqInfo)

	if err != nil {
		panic(err)
	}

	migrations.DbMigrate(DB)

	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET("/api/v1/country/", country_handler.GetAllCountries)
	router.GET("/api/v1/country/recipe/:countrycode", country_handler.GetAllCountriesRecipe)

	router.GET("/api/v1/recipe/", recipe_handler.GetAllRecipe)
	router.GET("/api/v1/recipe/:slug", recipe_handler.GetRecipeBySlug)
	router.GET("/api/v1/recipe/detail/:slug", recipe_handler.GetRecipeDetailBySlug)
	router.DELETE("/api/v1/recipe/:id", recipe_handler.DestroyRecipeById)

	router.DELETE("/api/v1/step/:id", recipe_handler.DeleteStepById)
	router.POST("/api/v1/step", recipe_handler.CreateStep)

	router.DELETE("/api/v1/ingredient/:id", recipe_handler.DeleteIngredientById)
	router.POST("/api/v1/ingredient", recipe_handler.CreateIngredient)

	router.GET("/api/v1/content/:recipes_id", content_handler.GetRecipeContent)

	router.POST("/api/v1/like", content_handler.CreateLike)
	router.DELETE("/api/v1/like/:id", content_handler.DestroyLikeById)

	router.POST("/api/v1/comment", content_handler.CreateComment)
	router.DELETE("/api/v1/comment/:id", content_handler.DestroyCommentById)

	router.GET("/api/v1/user/", user_handler.GetAllUser)
	router.GET("/api/v1/user/:slug", user_handler.GetUserBySlug)
	router.DELETE("/api/v1/user/:id", user_handler.DeleteUserById)

	router.Run("localhost:8080")
}
