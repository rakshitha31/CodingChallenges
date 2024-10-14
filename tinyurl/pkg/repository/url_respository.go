// create a connection to mongodb database running locally from dockerfile and return the connection

package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/rakshitha31/urlshortnerchallenge/pkg/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectToDB() *mongo.Client {
	username := "root"
	password := "P@assw0rd"
	escapedPassword := url.QueryEscape(password)
	// Construct the connection URI
	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017", username, escapedPassword)
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func DisconnectFromDB(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Disconnect(ctx)
	fmt.Println("Connection to MongoDB closed.")
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("urlshortner").Collection(collectionName)
}

func InsertOneDocument(collection *mongo.Collection, filter bson.M, document *model.Url) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := FindOneDocument(collection, filter)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		fmt.Println("Document already exists!")
		return
	}
	_, err = collection.InsertOne(ctx, document)
	if err != nil {
		log.Fatal(err)
	}
}

func FindOneDocument(collection *mongo.Collection, filter bson.M) (*model.Url, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result := new(model.Url)
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
