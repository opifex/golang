package app

type BookDeleteMessage struct {
	UUID string `form:"uuid"`
}

type BookDeleteResponse struct {
}

func (m BookDeleteMessage) Handle(core *Core) BookDeleteResponse {
	book, err := core.Repositories.Book.GetOneById(m.UUID)

	if err != nil {
		panic("book delete handle: " + err.Error())
	}

	if _, err := core.Repositories.Book.Delete(book); err != nil {
		panic("book create handle: " + err.Error())
	}

	return BookDeleteResponse{}
}
