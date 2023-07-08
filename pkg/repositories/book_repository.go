package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	guuid "github.com/google/uuid"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/thifnmi/go-book-api/pkg/domain"
	base_query "github.com/thifnmi/go-book-api/pkg/interfaces/http/query"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll(ctx context.Context, query *domain.BookQuery) ([]domain.BaseBookResponse, *domain.MetadataResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Book, error)
	CreateBook(ctx context.Context, payload *domain.BookPayload) (domain.Response, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewbookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		db,
	}
}

func (lr *bookRepository) GetAll(ctx context.Context, query *domain.BookQuery) ([]domain.BaseBookResponse, *domain.MetadataResponse, error) {
	var (
		book     []domain.Book
		basebook []domain.BaseBookResponse
		metaData domain.MetadataResponse
	)
	q := lr.db.Model(domain.Book{})

	metaData, err := base_query.GetPaginatedResults(ctx, lr.db, &book, q, query.Page, query.Limit)
	if err != nil {
		log.Printf("error %v when query metadata %v", err, query)
	}
	result := q.Limit(query.Limit).Find(&basebook)
	if result.Error != nil {
		log.Printf("query all with filter %v have err %v", query, result.Error)
		return nil, nil, result.Error
	}
	return basebook, &metaData, nil
}

func (lr *bookRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Book, error) {
	var book domain.Book
	lr.db.Where("uuid = ?", id).First(&book)
	return book, nil
}

func (lr *bookRepository) CreateBook(ctx context.Context, payload *domain.BookPayload) (domain.Response, error) {
	var response domain.Response
	ui := guuid.New()
	layout := "2006-01-02 15:04:05"
	t := time.Now().Format("2006-01-02 15:04:05")
	parsedTime, err := time.Parse(layout, t)
	if err != nil {
		response.Message = ""
		return response, errors.New(fmt.Sprintf("parse time to save log from consumer has err %v", err))
	}
	book := &domain.Book{
		Uuid:        ui,
		Name:        payload.Name,
		Category_id: payload.Category_id,
		Price:       payload.Price,
		CreatedAt:   &parsedTime,
	}

	result := lr.db.Create(book)

	if result.Error != nil {
		log.Panicf("failed to insert logs %v", result.Error)
		response.Message = ""
		return response, result.Error
	}
	return response, nil
}
