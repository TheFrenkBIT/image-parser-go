package models

import "time"

type Image struct {
	ImagePath    string
	ParsedResult string
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
