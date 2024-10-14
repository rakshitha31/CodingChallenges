package service

import (
	"fmt"

	"github.com/rakshitha31/urlshortnerchallenge/pkg/model"
	"github.com/rakshitha31/urlshortnerchallenge/pkg/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func AddUrl(longUrl, shortUrl, key string) model.Url {
	client := repository.ConnectToDB()
	defer repository.DisconnectFromDB(client)
	collection := repository.GetCollection(client, "urls")
	document := &model.Url{LongUrl: longUrl, ShortUrl: shortUrl, Key: key}
	filter := bson.M{"longurl": longUrl, "key": key}
	repository.InsertOneDocument(collection, filter, document)
	return *document

}

func GetLongUrl(key string) (string, error) {
	client := repository.ConnectToDB()
	defer repository.DisconnectFromDB(client)
	collection := repository.GetCollection(client, "urls")
	filter := bson.M{"key": key}
	UrlI, err := repository.FindOneDocument(collection, filter)
	if err != nil {
		fmt.Println("Error finding document")
		return "", err
	}
	var Url model.Url
	bsonBytes, _ := bson.Marshal(UrlI)
	bson.Unmarshal(bsonBytes, &Url)
	fmt.Println(Url.GetLongUrl())
	return Url.GetLongUrl(), nil

}
