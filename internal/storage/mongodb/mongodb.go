package mongodb

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// type MongoDB struct {
// 	client *mongo.Database
// }

func FindOneDoc(db *mongo.Database, ctx context.Context, collection string, data map[string]interface{}) (interface{}, error) {
	collectionRef := db.Collection(collection)
	idStr, ok := data["_id"].(string)
	if !ok {
		return nil, fmt.Errorf("_id is not a string")
	}
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID: %v", err)
	}
	filter := bson.M{"_id": id}
	err = collectionRef.FindOne(ctx, filter).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func InsertOneDoc(db *mongo.Database, ctx context.Context, collection string, data map[string]interface{}) (interface{}, error) {
	collectionRef := db.Collection(collection)
	val, err := collectionRef.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	data["_id"] = val.InsertedID
	return data, nil
}

func FindAllDoc(db *mongo.Database, ctx context.Context, collection string) (map[string]interface{}, error) {
	collectionRef := db.Collection(collection)
	cur, err := collectionRef.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var data []map[string]interface{}
	for cur.Next(ctx) {
		var item map[string]interface{}
		err := cur.Decode(&item)
		fmt.Println(item)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}
	// count total doc
	count, err := collectionRef.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cur.Err()
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	val := map[string]interface{}{
		"total": count,
		"data":  data,
	}
	return val, nil
}

func UpdateDoc(db *mongo.Database, ctx context.Context, collection string, id string, data map[string]interface{}) (map[string]interface{}, error) {
	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	// Create the filter and update documents
	filter := bson.D{{"_id", objectID}}
	update := bson.D{{"$set", data}}
	fmt.Println("Filter:", filter)
	fmt.Println("Update:", update)

	// Get the collection reference
	collectionRef := db.Collection(collection)

	// Define options to return the updated document after update
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	// Perform the update and retrieve the updated document
	var updatedDocument map[string]interface{}
	err = collectionRef.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedDocument)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no document found with the provided ID")
		}
		return nil, fmt.Errorf("failed to update document: %v", err)
	}

	return updatedDocument, nil
}

// type mongodb interface {
// 	InsertOne(ctx context.Context, collection string, data interface{}) error
// 	FindOne(ctx context.Context, collection string, filter interface{}) (*mongo.SingleResult, error)
// 	Find(ctx context.Context, collection string, filter interface{}) (interface{}, error)
// 	UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
// 	DeleteOne(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)
// }
