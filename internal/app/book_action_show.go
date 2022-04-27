package app

import "time"

type BookShowMessage struct {
	UUID string `form:"uuid"`
}

type BookShowResponse struct {
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m BookShowMessage) Handle(core *Core) BookShowResponse {
	document, err := core.Repositories.Book.GetOneById(m.UUID)

	if err != nil {
		panic("book show handle: " + err.Error())
	}

	return BookShowResponse{
		Author:    document.Author,
		CreatedAt: document.CreatedAt,
		Title:     document.Title,
		UUID:      document.UUID.String(),
		UpdatedAt: document.UpdatedAt,
		Year:      document.Year,
	}
}
