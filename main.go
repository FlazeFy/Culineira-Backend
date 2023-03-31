package main

import (
	"culineira-backend/middlewares"
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

	// Public API
	public := router.Group("/api")
	public.POST("/v1/login", auth_handler.Login)
	public.POST("/v1/register", auth_handler.Register)

	public.GET("/v1/country/", country_handler.GetAllCountries)
	public.GET("/v1/country/recipe/:countrycode", country_handler.GetAllCountriesRecipe)
	public.GET("/v1/recipe/:slug", recipe_handler.GetRecipeBySlug)

	// Protected JWT API
	protected := router.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	protected.GET("/v1/recipe/", recipe_handler.GetAllRecipe)
	protected.GET("/v1/recipe/detail/:slug", recipe_handler.GetRecipeDetailBySlug)
	protected.POST("/v1/recipe", recipe_handler.CreateRecipe)
	protected.DELETE("/v1/recipe/:id", recipe_handler.DestroyRecipeById)
	protected.PUT("/v1/recipe/:id", recipe_handler.UpdateRecipeById)

	protected.DELETE("/v1/step/:id", recipe_handler.DeleteStepById)
	protected.POST("/v1/step", recipe_handler.CreateStep)
	protected.PUT("/v1/step/:id", recipe_handler.UpdateStepById)

	protected.DELETE("/v1/ingredient/:id", recipe_handler.DeleteIngredientById)
	protected.POST("/v1/ingredient", recipe_handler.CreateIngredient)
	protected.PUT("/v1/ingredient/:id", recipe_handler.UpdateIngredientById)

	protected.GET("/v1/content/:recipes_id", content_handler.GetRecipeContent)

	protected.POST("/v1/like", content_handler.CreateLike)
	protected.DELETE("/v1/like/:id", content_handler.DestroyLikeById)

	protected.POST("/v1/comment", content_handler.CreateComment)
	protected.DELETE("/v1/comment/:id", content_handler.DestroyCommentById)

	protected.GET("/v1/user/", user_handler.GetAllUser)
	protected.GET("/v1/user/:slug", user_handler.GetUserBySlug)
	protected.DELETE("/v1/user/:id", user_handler.DeleteUserById)
	protected.PUT("/v1/user/:id", user_handler.UpdateUserById)

	router.Run(":" + os.Getenv("DB_PORT"))
}
