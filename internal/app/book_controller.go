package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	Core *Core
}

func (c *BookController) Create(ctx *gin.Context) {
	var message = c.Core.Messages.BookCreate
	if err := ctx.BindJSON(&message); err != nil {
		return
	}
	response := message.Handle(c.Core)
	ctx.JSON(http.StatusCreated, response)
}

func (c *BookController) Delete(ctx *gin.Context) {
	var message = c.Core.Messages.BookDelete
	message.UUID = ctx.Param("uuid")
	response := message.Handle(c.Core)
	ctx.JSON(http.StatusNoContent, response)
}

func (c *BookController) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Book Index")
}

func (c *BookController) Show(ctx *gin.Context) {
	var message = c.Core.Messages.BookShow
	message.UUID = ctx.Param("uuid")
	response := message.Handle(c.Core)
	ctx.JSON(http.StatusOK, response)
}
