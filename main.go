package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Garry028/mongoApi/database"
	"github.com/Garry028/mongoApi/router"
	"github.com/joho/godotenv"
	// Import the godotenv package
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read environment variables
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	databaseName := os.Getenv("DB_NAME")

	// Initialize the database connection
	database.InitDatabase(connectionString, databaseName)
	fmt.Println("Server is getting started...")
	r := router.SetupRoutes()

	log.Fatal(http.ListenAndServe(":4000", r))
}
