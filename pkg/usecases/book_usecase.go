package usecases

import (
	"context"
	"errors"
	"fmt"

	"github.com/satori/go.uuid"
	"github.com/thifnmi/go-book-api/pkg/domain"
	"github.com/thifnmi/go-book-api/pkg/repositories"
)

type BookUsecase interface {
	GetAll(ctx context.Context, query *domain.BookQuery) (domain.ListBookResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.BookResponse, error)
	CreateBook(ctx context.Context, payload *domain.BookPayload) (domain.Response, error)
}

type bookUsecase struct {
	bookRepo repositories.BookRepository
}

func NewBookUsecase(bookRepo repositories.BookRepository) BookUsecase {
	return &bookUsecase{
		bookRepo,
	}
}

func (lu *bookUsecase) GetAll(ctx context.Context, query *domain.BookQuery) (domain.ListBookResponse, error) {
	books, metaData, err := lu.bookRepo.GetAll(ctx, query)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	response := domain.ListBookResponse{
		Success:   true,
		ErrorCode: 0,
		Message:   "List books success",
		Data:      make([]domain.BaseBookResponse, 0),
	}
	if metaData != nil {
		response.Metadata = domain.MetadataResponse{
			Total:       metaData.Total,
			CurrentPage: metaData.CurrentPage,
			Pages:       metaData.Pages,
		}
	}
	for _, book := range books {
		baseBook := domain.BaseBookResponse{
			Uuid:        book.Uuid,
			Name:        book.Name,
			Category_id: book.Category_id,
			Price:       book.Price,
			CreatedAt:   book.CreatedAt,
		}

		response.Data = append(response.Data, baseBook)
	}
	return response, err
}

func (lu *bookUsecase) GetByID(ctx context.Context, id uuid.UUID) (domain.BookResponse, error) {
	book, err := lu.bookRepo.GetByID(ctx, id)

	response := domain.BookResponse{
		Success:   true,
		ErrorCode: 0,
		Message:   "Get Book success",
	}

	baseBook := domain.BaseBookResponse{
		Uuid:        book.Uuid,
		Name:        book.Name,
		Category_id: book.Category_id,
		Price:       book.Price,
		CreatedAt:   book.CreatedAt,
	}
	response.Data = baseBook
	return response, err
}

func (lr *bookUsecase) CreateBook(ctx context.Context, payload *domain.BookPayload) (domain.Response, error) {
	response, err := lr.bookRepo.CreateBook(ctx, payload)
	if err != nil {
		return response, errors.New(fmt.Sprintf("Failed to create book with err %v", err))
	}
	return response, nil
}
