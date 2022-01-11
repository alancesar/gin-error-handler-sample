package database

import (
	"context"
	"fmt"
	"github.com/alancesar/gin-error-handler-sample/pkg"
)

type (
	Database struct{}

	Err struct {
		Reason string
		Code   string
	}
)

func (e Err) Error() string {
	return fmt.Sprintf("%s (%s)", e.Reason, e.Code)
}

func (e Err) Is(target error) bool {
	return target == pkg.ErrInternal
}

func (d Database) GetCustomer(_ context.Context, fail bool) error {
	if fail {
		return newErr("failed on database", "DB_ERROR")
	}

	return nil
}

func newErr(reason, code string) *Err {
	return &Err{
		Reason: reason,
		Code:   code,
	}
}

func NewDatabase() *Database {
	return &Database{}
}
