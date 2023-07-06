package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thifnmi/go-book-api/pkg/domain"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Heathcheck(c *gin.Context) {
	c.JSON(
		http.StatusOK, &domain.HealthResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Message:    "healthy",
		},
	)
}

func (h *HealthHandler) Info(c *gin.Context) {
	c.JSON(
		http.StatusOK, &domain.InfoResponse{
			Auth:     "Thifnmi",
			FullName: "Thin Tu Van",
			Email:    "tuthin2k@gmail.com",
			Telegram: "@Thifnmi",
		},
	)
}
