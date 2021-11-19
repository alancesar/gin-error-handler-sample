package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	request struct {
		ID           string `json:"id" binding:"required"`
		FailService  bool   `json:"fail_service"`
		FailDatabase bool   `json:"fail_database"`
	}

	service interface {
		GetCustomer(ctx context.Context, failService, failDatabase bool) error
	}
)

func Handler(s service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		// Using ShouldBindJSON will not set Content-Type header with text/plain; charset=utf-8
		// https://github.com/gin-gonic/gin#model-binding-and-validation
		if err := c.ShouldBindJSON(&req); err != nil {
			_ = c.Error(err)
			return
		}

		if err := s.GetCustomer(c.Request.Context(), req.FailService, req.FailDatabase); err != nil {
			_ = c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":   req.ID,
			"name": "Error Handler",
		})
	}
}
