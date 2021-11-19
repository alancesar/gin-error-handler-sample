package service

import (
	"context"
	"github.com/alancesar/gin-error-handler-sample/pkg"
)

type (
	database interface {
		GetCustomer(ctx context.Context, fail bool) error
	}

	Err struct {
		Message string
		Code    int
	}

	Service struct {
		db database
	}
)

func (e Err) Error() string {
	return e.Message
}

func (e Err) Is(target error) bool {
	return target == pkg.InternalErr
}

func (i Service) GetCustomer(ctx context.Context, failService, failDatabase bool) error {
	if failService {
		return newErr("failed on service", 500)
	}

	return i.db.GetCustomer(ctx, failDatabase)
}

func newErr(message string, code int) *Err {
	return &Err{
		Message: message,
		Code:    code,
	}
}

func NewService(db database) *Service {
	return &Service{
		db: db,
	}
}
