package entity

import (
	"errors"
)

var (
	ErrInternal        = errors.New("internal error")
	ErrDataNotFound    = errors.New("data not found")
	ErrConflictingData = errors.New("data conflicts with existing data in unique column")
	ErrForbidden       = errors.New("user is forbidden to access the resource")
	ErrNoUpdatedData   = errors.New("no data to update")
)
