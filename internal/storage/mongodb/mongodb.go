package mongodb

import (
	"context"
	"log"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb(uri string) *mongo.Database {
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("error in mongodb connection: ", err)
	}
	// Ping the server to check the connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("db is not healthy : ", err)
	}
	slog.Info("MongoDB connected successfully!!")
	return client.Database("movie")
}
