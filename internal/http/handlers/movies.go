package movies

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func WelcomeFun(storage *mongo.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Server status checked : ", slog.String("ip", r.RemoteAddr))
		// insert in mogo
		collection := storage.Collection("server_status")
		result, err := collection.InsertOne(context.TODO(), bson.M{"ip": r.RemoteAddr})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "welcome to movie api",
			"success": true,
			"result":  result,
		})
	}
}
