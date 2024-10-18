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
	routers.HandleFunc("/getMovies", movies.GetMovies(storage)).Methods("GET")
	routers.HandleFunc("/createMovie", movies.CreateMovie(storage)).Methods("POST")
	routers.HandleFunc("/getById/{id}", movies.GetMovieById(storage)).Methods("GET")
	return routers
}
