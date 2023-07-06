package query

import (
	"context"
	"math"

	"github.com/thifnmi/go-book-api/pkg/domain"
	"gorm.io/gorm"
)

type PaginationMetadata struct {
	CurrentPage  int  `json:"current_page"`
	Page         int  `json:"page"`
	TotalPages   int  `json:"total_pages"`
	HasPrevious  bool `json:"has_previous"`
	HasNext      bool `json:"has_next"`
	PreviousPage int  `json:"previous_page"`
	NextPage     int  `json:"next_page"`
}

func GetPaginatedResults(ctx context.Context, db *gorm.DB, results interface{}, query *gorm.DB, page, limit int) (domain.MetadataResponse, error) {
	var count int64
	err := db.Model(results).Where(query).Count(&count).Error
	if err != nil {
		return domain.MetadataResponse{}, err
	}

	var metadata domain.MetadataResponse
	metadata.Total = int(count)
	metadata.CurrentPage = page

	if limit < 1 {
		metadata.Pages = 1
	} else {
		metadata.Pages = int(math.Ceil(float64(count) / float64(limit)))
	}
	return metadata, nil
}
