package Database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Friend struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID         string      `json:"id,omiempty"`
	Password   string      `json:"password,omiempty"`
	Isactive   bool        `json:"isactive,omiempty"`
	Balance    string      `json:"balance,omiempty"`
	Age        json.Number `json:"age,omiempty"`
	Name       string      `json:"name,omiempty"`
	Gender     string      `json:"gender,omiempty"`
	Company    string      `json:"company,omiempty"`
	Email      string      `json:"email,omiempty"`
	Phone      string      `json:"phone,omiempty"`
	Address    string      `json:"address,omiempty"`
	About      string      `json:"about,omiempty"`
	Registered string      `json:"registered,omiempty"`
	Latitude   float32     `json:"latitude,omiempty"`
	Longitude  float32     `json:"longitude,omiempty"`
	Tags       []string    `json:"tags,omiempty"`
	Friends    []Friend    `json:"friends,omiempty"`
	Data       string      `json:"data,omiempty"`
}

var client_mongo *mongo.Client

func ConnectDB(url string) error {
	client_mongo, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return err
	}

	ctx, cancelTimeout := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelTimeout()
	err = client_mongo.Connect(ctx)
	if err != nil {
		return err
	}

	/* Pour tester la connection à Mongo */
	err = client_mongo.Ping(ctx, nil)
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB")
	return nil
}

func getUsersCollection(collectionName string) *mongo.Collection {
	collection := client_mongo.Database("users_db").Collection(collectionName)
	return collection
}

func DeleteOne(ctx context.Context, bson bson.M) (*mongo.DeleteResult, error) {
	return getUsersCollection("users_db").DeleteOne(ctx, bson)
}

func FindOne(ctx context.Context, bson bson.M) *mongo.SingleResult {
	return getUsersCollection("users_db").FindOne(ctx, bson)
}

func Find(ctx context.Context, bson bson.M) (*mongo.Cursor, error) {
	return getUsersCollection("users_db").Find(ctx, bson)
}

func UpdateOne(ctx context.Context, bson bson.M, bson_set bson.M, options *options.UpdateOptions) (*mongo.UpdateResult, error) {
	return getUsersCollection("users_db").UpdateOne(ctx, bson, bson_set, options)
}

func FindOneAndUpdate(ctx context.Context, bson bson.M, bson_set bson.M) *mongo.SingleResult {
	return getUsersCollection("users_db").FindOneAndUpdate(ctx, bson, bson_set)
}
