package main

import (
	"culineira-backend/migrations"
	auth_handler "culineira-backend/modules/auth/handlers"
	content_handler "culineira-backend/modules/contents/handlers"
	country_handler "culineira-backend/modules/countries/handlers"
	recipe_handler "culineira-backend/modules/recipes/handlers"
	user_handler "culineira-backend/modules/users/handlers"
	"database/sql"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Failed to load env")
		panic(err)
	} else {
		fmt.Println("Success to load env")
	}

	psqInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

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
	router.POST("/api/v1/recipe", recipe_handler.CreateRecipe)
	router.DELETE("/api/v1/recipe/:id", recipe_handler.DestroyRecipeById)
	router.PUT("/api/v1/recipe/:id", recipe_handler.UpdateRecipeById)

	router.DELETE("/api/v1/step/:id", recipe_handler.DeleteStepById)
	router.POST("/api/v1/step", recipe_handler.CreateStep)
	router.PUT("/api/v1/step/:id", recipe_handler.UpdateStepById)

	router.DELETE("/api/v1/ingredient/:id", recipe_handler.DeleteIngredientById)
	router.POST("/api/v1/ingredient", recipe_handler.CreateIngredient)
	router.PUT("/api/v1/ingredient/:id", recipe_handler.UpdateIngredientById)

	router.GET("/api/v1/content/:recipes_id", content_handler.GetRecipeContent)

	router.POST("/api/v1/like", content_handler.CreateLike)
	router.DELETE("/api/v1/like/:id", content_handler.DestroyLikeById)

	router.POST("/api/v1/comment", content_handler.CreateComment)
	router.DELETE("/api/v1/comment/:id", content_handler.DestroyCommentById)

	router.GET("/api/v1/user/", user_handler.GetAllUser)
	router.GET("/api/v1/user/:slug", user_handler.GetUserBySlug)
	router.DELETE("/api/v1/user/:id", user_handler.DeleteUserById)
	router.PUT("/api/v1/user/:id", user_handler.UpdateUserById)

	router.POST("/api/v1/login", auth_handler.Login)

	router.Run("localhost:8080")
}
