package domain

import (
	"time"

	uuid "github.com/google/uuid"
)

type Book struct {
	ID          int        `gorm:"primaryKey"`
	Uuid        uuid.UUID  `json:"uuid" gorm:"not null"`
	Name        string     `json:"name"`
	Category_id uuid.UUID  `json:"category_id"`
	Price       int        `json:"price"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type BaseBookResponse struct {
	Uuid        uuid.UUID  `json:"uuid"`
	Name        string     `json:"name"`
	Category_id uuid.UUID  `json:"category_id"`
	Price       int        `json:"price"`
	CreatedAt   *time.Time `json:"created_at"`
}

type BookResponse struct {
	Success   bool             `json:"success"`
	ErrorCode int16            `json:"error_code"`
	Message   string           `json:"message"`
	Data      BaseBookResponse `json:"data"`
}

type ListBookResponse struct {
	Success   bool               `json:"success"`
	ErrorCode int16              `json:"error_code"`
	Message   string             `json:"message"`
	Data      []BaseBookResponse `json:"data"`
	Metadata  MetadataResponse   `json:"metadata"`
}

type BookQuery struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type BookPayload struct {
	Name        string    `form:"name"`
	Category_id uuid.UUID `form:"category_id"`
	Price       int       `form:"price"`
}
