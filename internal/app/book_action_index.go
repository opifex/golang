package app

type BookIndexMessage struct {
	UUID string `form:"uuid"`
}

type BookIndexResponse struct {
}

func (m BookIndexMessage) Handle(core *Core) BookIndexResponse {
	return BookIndexResponse{}
}
