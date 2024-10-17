package router

import (
	movies "github.com/amit8889/go-movie-api/internal/http/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func MovieRouter(storage *mongo.Database) *mux.Router {
	// Define the routes
	routers := mux.NewRouter()
	routers.HandleFunc("/", movies.WelcomeFun(storage)).Methods("GET")
	return routers
}
