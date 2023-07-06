package book

import (
	"github.com/gin-gonic/gin"
	"github.com/thifnmi/go-book-api/pkg/interfaces/http/handlers"
)

func SetupRouter(router *gin.Engine, bookHandler handlers.BookHandler) {
	api := router.Group("/books")
	{
		api.GET("", bookHandler.GetAll)
		api.GET("/:uuid", bookHandler.GetByID)
	}
}
