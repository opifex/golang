package app

import (
	"time"

	"github.com/google/uuid"
)

type BookCreateMessage struct {
	Author string `form:"author" valid:"Required;MaxSize(50)"`
	Title  string `form:"title" valid:"Required;MaxSize(50)"`
	Year   int    `form:"year" valid:"Required;Min(0)"`
}

type BookCreateResponse struct {
	Identifier string `json:"identifier"`
}

func (m BookCreateMessage) Handle(core *Core) BookCreateResponse {
	book := Book{
		Author:    m.Author,
		CreatedAt: time.Now(),
		Title:     m.Title,
		UUID:      uuid.New(),
		UpdatedAt: time.Now(),
		Year:      m.Year,
	}

	if _, err := core.Repositories.Book.Create(book); err != nil {
		panic("book create handle: " + err.Error())
	}

	return BookCreateResponse{Identifier: book.UUID.String()}
}
