package router

import (
	"net/http"

	"github.com/Garry028/mongoApi/controller"
)

func SetupRoutes() {
	http.HandleFunc("/movies", controller.GetMyAllMovies)
	// Set up other routes
}
