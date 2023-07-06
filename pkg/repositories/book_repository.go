package repositories

import (
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/thifnmi/go-book-api/pkg/domain"
	base_query "github.com/thifnmi/go-book-api/pkg/interfaces/http/query"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll(ctx context.Context, query *domain.BookQuery) ([]domain.BaseBookResponse, *domain.MetadataResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Book, error)
	CreateBook(ctx context.Context, payload *domain.BookPayload) (domain.Book, error)
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
		fmt.Sprintf("error %v when query metadata %v", err, query)
	}
	// span, ctx = apm.StartSpan(ctx, "GetAll Repository", "DB MySQL query results")
	// defer span.End()
	result := q.Limit(query.Limit).Offset(10).Find(&basebook)
	if result.Error != nil {
		fmt.Sprintf("query all with filter %v have err %v", query, result.Error)
		return nil, nil, result.Error
	}
	return basebook, &metaData, nil
}

func (lr *bookRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Book, error) {
	var book domain.Book
	lr.db.Where("uuid = ?", id).First(&book)
	return book, nil
}

func (lr *bookRepository) CreateBook(ctx context.Context, payload *domain.BookPayload) (domain.Book, error) {
	var book domain.Book
	return book, nil
}
