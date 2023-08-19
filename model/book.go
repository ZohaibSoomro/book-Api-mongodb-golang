package model

import (
	"context"
	"fmt"

	"github.com/zohaibsoomro/book-server-mongodb/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Author      string             `json:"author" bson:"author"`
	PublishDate string             `json:"publish_date" bson:"publish_date"`
}

var client *mongo.Client

func SetClient(cl *mongo.Client) {
	client = cl
}

func (b *Book) CreateBookInDB() error {
	b.Id = primitive.NewObjectID()
	res, err := client.Database(config.DbName).Collection(config.CollectionName).InsertOne(context.Background(), b)
	if err != nil {
		return err
	}
	fmt.Printf("Book created with id %v", res.InsertedID)
	return nil
}

func GetBookWithIdFromDB(Id primitive.ObjectID) (*Book, error) {
	b := &Book{}
	err := client.Database(config.DbName).Collection(config.CollectionName).FindOne(context.Background(), bson.D{{Key: "_id", Value: Id}}).Decode(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetAllBooksFromDb() ([]Book, error) {
	var books []Book
	cur, err := client.Database(config.DbName).Collection(config.CollectionName).Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		b := Book{}
		err := cur.Decode(&b)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (book *Book) DeleteBookWithIdFromDb() (*Book, error) {
	b, err := GetBookWithIdFromDB(book.Id)
	//if book not exists with id
	if err != nil {
		return nil, err
	}
	// next delete book
	res, err := client.Database(config.DbName).Collection(config.CollectionName).DeleteOne(context.Background(), bson.D{{Key: "_id", Value: book.Id}})
	if err != nil {
		return nil, err
	}
	fmt.Println("Total deleted records:", res.DeletedCount)
	return b, nil
}

func (b *Book) UpdateBookInDb() error {
	res, err := client.Database(config.DbName).Collection(config.CollectionName).UpdateByID(context.Background(), bson.D{{Key: "_id", Value: b.Id}}, b)
	if err != nil {
		return err
	}
	fmt.Println("Total deleted records:", res.UpsertedCount)
	return nil
}
