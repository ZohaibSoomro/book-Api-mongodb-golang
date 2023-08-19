package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DbName = "booksdb"
const CollectionName = "books"

func ConnectDB() *mongo.Client {
	dns := "mongodb+srv://zhs:Zohaib123@cluster0.8oxz7j7.mongodb.net/?retryWrites=true&w=majority"
	_client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dns))
	if err != nil {
		panic(err)
	}
	return _client
}

func PingDB(client *mongo.Client) error {

	if err := client.Ping(context.Background(), nil); err != nil {
		return err
	}
	return nil
}
