package app

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
)

type Configuration struct {
	Connection *MongoAdapter
}

type Core struct {
	context.Context
	Messages     Messages
	Repositories Repositories
}

type Messages struct {
	BookCreate BookCreateMessage
	BookDelete BookDeleteMessage
	BookIndex  BookIndexMessage
	BookShow   BookShowMessage
}

type Repositories struct {
	Book *BookRepository
}

func Init(configuration Configuration) *Core {
	return &Core{
		Repositories: Repositories{
			Book: NewBookRepository(configuration.Connection),
		},
	}
}

func Run() {
	gin.SetMode(os.Getenv("GIN_MODE"))

	connection, err := MongoConnection(os.Getenv("MONGODB_URL"), os.Getenv("MONGODB_DB"))

	if err != nil {
		panic("error connecting to database")
	}

	core := Init(Configuration{Connection: connection})

	HandlePortHttp(core)

	if err = connection.Disconnect(); err != nil {
		return
	}
}
