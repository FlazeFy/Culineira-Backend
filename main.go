package main

import (
	"culineira-backend/migrations"
	"culineira-backend/modules/countries/handlers"
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
	err = DB.Ping()

	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	migrations.DbMigrate(DB)

	defer DB.Close()

	router := gin.Default()

	router.GET("/api/v1/country/", handlers.GetAllCountries)
	router.GET("/api/v1/country/recipe/", handlers.GetAllCountriesRecipe)

	router.Run("localhost:8080")
}
