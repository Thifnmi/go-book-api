package handlers

import (
	"fmt"
	"log"
	"net/http"
	// "io/ioutil"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/thifnmi/go-book-api/pkg/domain"
	"github.com/thifnmi/go-book-api/pkg/usecases"
)

type BookHandler interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	CreateBook(c *gin.Context)
}

type bookHandler struct {
	bookUsecase usecases.BookUsecase
}

func NewbookHandler(bookUsecase usecases.BookUsecase) BookHandler {
	return &bookHandler{
		bookUsecase,
	}
}

func (lh *bookHandler) GetAll(c *gin.Context) {
	var query domain.BookQuery
	err := c.BindQuery(&query)
	if err != nil {
		fmt.Printf("err %v'n", err)
	}
	books, err := lh.bookUsecase.GetAll(c, &query)
	if err != nil {
		fmt.Printf("err %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (lh *bookHandler) GetByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("uuid"))
	if err != nil {
		fmt.Printf("err %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := lh.bookUsecase.GetByID(c, id)
	if err != nil {
		fmt.Printf("err %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (lh *bookHandler) CreateBook(c *gin.Context) {
	bookRequestBody := &domain.BookPayload{}
	if err := c.BindJSON(&bookRequestBody); err != nil {
		log.Print("Can't bind body request")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return 
	}
	response, err := lh.bookUsecase.CreateBook(c, bookRequestBody)
	if err != nil {
		fmt.Printf("err %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Print(bookRequestBody)
	c.JSON(http.StatusOK,response)
}
