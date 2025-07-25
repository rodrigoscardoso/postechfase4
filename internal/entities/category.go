package entity

import (
	"time"
)

type Category struct {
	Id        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
