package app

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "book"

type BookRepository struct {
	adapter    *MongoAdapter
	collection *mongo.Collection
}

func NewBookRepository(adapter *MongoAdapter) *BookRepository {
	return &BookRepository{
		adapter:    adapter,
		collection: adapter.AssignCollection(collectionName),
	}
}

func (r *BookRepository) Create(book Book) (interface{}, error) {
	return r.adapter.CreateDocument(r.collection, book)
}

func (r *BookRepository) GetOneById(identifier string) (Book, error) {
	book := Book{}

	if _, err := r.adapter.GetDocument(r.collection, uuid.MustParse(identifier), &book); err != nil {
		return book, err
	}

	return book, nil
}

func (r *BookRepository) Delete(book Book) (interface{}, error) {
	return r.adapter.DeleteDocument(r.collection, book)
}
