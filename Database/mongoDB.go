package Database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

type DatabaseHelper interface {
	Collection(name string) CollectionHelper
}

type CollectionHelper interface {
	Find(context.Context, interface{}, ...*options.FindOptions) (CursorHelper, error)
	FindOne(context.Context, interface{}, ...*options.FindOneOptions) SingleResultHelper
	UpdateOne(context.Context, interface{}, interface{}, *options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(context.Context, interface{}, ...*options.DeleteOptions) (int64, error)
}

type SingleResultHelper interface {
	Decode(interface{}) error
}

type CursorHelper interface {
	Decode(interface{}) error
	Next(context.Context) bool
	Close(context.Context) error
}

type ClientHelper interface {
	Database(string) DatabaseHelper
	Connect(context.Context) error
	Ping(context.Context, *readpref.ReadPref) error
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoCursor struct {
	cs *mongo.Cursor
}

func newClient(url string) (ClientHelper, error) {
	c, err := mongo.NewClient(options.Client().ApplyURI(url))
	return &mongoClient{cl: c}, err
}

func (mc *mongoClient) Database(dbName string) DatabaseHelper {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) Connect(ctx context.Context) error {
	return mc.cl.Connect(ctx)
}

func (mc *mongoClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return mc.cl.Ping(ctx, rp)
}

func (md *mongoDatabase) Collection(colName string) CollectionHelper {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}, options ...*options.FindOneOptions) SingleResultHelper {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) Find(ctx context.Context, document interface{}, options ...*options.FindOptions) (CursorHelper, error) {
	cursor, err := mc.coll.Find(ctx, document)
	return cursor, err
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}, options ...*options.DeleteOptions) (int64, error) {
	count, err := mc.coll.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

func (mc *mongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, options *options.UpdateOptions) (*mongo.UpdateResult, error) {
	updateResult, err := mc.coll.UpdateOne(ctx, filter, update, options)
	return updateResult, err
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}

func (cursor *mongoCursor) Decode(v interface{}) error {
	return cursor.cs.Decode(v)
}

func (cursor *mongoCursor) Next(ctx context.Context) bool {
	return cursor.cs.Next(ctx)
}

func (cursor *mongoCursor) Close(ctx context.Context) error {
	return cursor.cs.Close(ctx)
}

var client_mongo ClientHelper

func ConnectDB(url string) error {
	var err error
	client_mongo, err = newClient(url)
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

func GetUsersCollection(collectionName string) CollectionHelper {
	collection := client_mongo.Database("users_db").Collection(collectionName)
	return collection
}
