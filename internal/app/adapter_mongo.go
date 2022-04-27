package app

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAdapter struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func MongoConnection(uri string, db string) (*MongoAdapter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	mongoClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	return &MongoAdapter{Client: mongoClient, Database: mongoClient.Database(db)}, nil
}

func (c *MongoAdapter) Disconnect() error {
	err := c.Client.Disconnect(context.TODO())

	if err != nil {
		return err
	}

	fmt.Println("adapter to MongoDB closed.")

	return nil
}

func (c *MongoAdapter) AssignCollection(collectionName string) *mongo.Collection {
	return c.Database.Collection(collectionName)
}

func (c *MongoAdapter) GetDocument(collection *mongo.Collection, documentId uuid.UUID, model interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	findOptions := options.FindOne()
	filterInterface := bson.M{"uuid": documentId}

	return collection.FindOne(ctx, filterInterface, findOptions).Decode(model), nil
}

func (c *MongoAdapter) CreateDocument(collection *mongo.Collection, document interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, document)

	if err != nil {
		fmt.Println("error inserting document into database :: ", err)
		return "", err
	}

	return result.InsertedID, nil
}

func (c *MongoAdapter) DeleteDocument(collection *mongo.Collection, document interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, document)

	if err != nil {
		fmt.Println("error deleting document from database :: ", err)
		return "", err
	}

	return result.DeletedCount, nil
}
