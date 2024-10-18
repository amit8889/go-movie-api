package movies

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/amit8889/go-movie-api/internal/storage/mongodb"
	"github.com/amit8889/go-movie-api/internal/types"
	"github.com/amit8889/go-movie-api/internal/utils/response"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func WelcomeFun(storage *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Server status checked : ", slog.String("ip", r.RemoteAddr))
		// insert in mogo
		//collection := storage.Collection("server_status")
		// result, err := collection.InsertOne(context.TODO(), bson.M{"ip": r.RemoteAddr})
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		response.WriteResponse(w, http.StatusAccepted, map[string]interface{}{
			"message": "Server status checked",
			"success": true,
		})
	}
}
func GetMovieById(storage *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Get Movies called : ", slog.String("ip", r.RemoteAddr))
		vars := mux.Vars(r) // This extracts path variables
		id, ok := vars["id"]
		slog.Info("Id of user is : ", slog.String("id", id))
		if !ok || id == "" {
			http.Error(w, "id is required", http.StatusBadRequest)
			return
		}
		cur, err := mongodb.FindOneDoc(storage, context.TODO(), "moives", map[string]interface{}{"_id": id})
		if err != nil {
			response.WriteResponse(w, http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
				"success": false,
			})
			return
		}
		response.WriteResponse(w, http.StatusCreated, map[string]interface{}{
			"message": "Movie details!!",
			"success": true,
			"movie":   cur,
		})

	}
}
func CreateMovie(storage *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Create Movies called : ", slog.String("ip", r.RemoteAddr))
		var movie types.Movie
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			response.WriteResponse(w, http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
				"success": false,
			})
			return
		}
		validationErrors := response.ValidateStruct(movie)
		if validationErrors != nil {
			response.WriteResponse(w, http.StatusBadRequest, map[string]interface{}{
				"message": validationErrors,
				"success": false,
			})
			return
		}
		cur, err := mongodb.InsertOneDoc(storage, context.TODO(), "moives", map[string]interface{}{
			"title": movie.Title,
			"year":  movie.Year,
		})
		if err != nil {
			response.WriteResponse(w, http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
				"success": false,
			})
			return
		}
		response.WriteResponse(w, http.StatusCreated, map[string]interface{}{
			"message": "Movie created successfully",
			"success": true,
			"movie":   cur,
		})

	}
}

func GetMovies(storage *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Get Movies called : ", slog.String("ip", r.RemoteAddr))
		cur, err := mongodb.FindAllDoc(storage, context.TODO(), "moives")
		if err != nil {
			response.WriteResponse(w, http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
				"success": false,
			})
			return
		}

		response.WriteResponse(w, http.StatusOK, map[string]interface{}{
			"message": "Movies retrieved successfully",
			"success": true,
			"movies":  cur,
		})
	}
}
