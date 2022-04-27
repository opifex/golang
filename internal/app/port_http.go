package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

func HandlePortHttp(core *Core) {
	engine := gin.Default()

	bookController := BookController{Core: core}

	engine.PUT("/book", bookController.Create)
	engine.DELETE("/book/:uuid", bookController.Delete)
	engine.GET("/book", bookController.Index)
	engine.GET("/book/:uuid", bookController.Show)

	if err := engine.Run("0.0.0.0:" + os.Getenv("HTTP_PORT")); err != nil {
		panic("can not run port listener")
	}
}
