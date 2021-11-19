package middleware

import (
	"errors"
	"github.com/alancesar/gin-error-handler-sample/database"
	"github.com/alancesar/gin-error-handler-sample/pkg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if err := c.Errors.Last(); err != nil {
			// service.Err and database.Err implement Is(target error) bool returning true for pkg.InternalErr
			// so errors.Is return true for them
			// https://pkg.go.dev/errors#Is
			if errors.Is(err.Err, pkg.InternalErr) {
				log.Printf("internal server error: %s", err.Error())

				// Handle with specific error
				if dbErr, ok := err.Err.(*database.Err); ok {
					log.Printf("got a database error with code %s", dbErr.Code)
				}

				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "unexpected error. please try again later",
				})
				return
			}

			log.Printf("bad request error: %s", err.Error())
			c.JSON(http.StatusBadRequest, err)
		}
	}
}
