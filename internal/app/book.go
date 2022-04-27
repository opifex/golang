package app

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Author    string    `json:"author" bson:"author"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	Title     string    `json:"title" bson:"title"`
	UUID      uuid.UUID `json:"uuid" bson:"uuid"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	Year      int       `json:"year" bson:"year"`
}
