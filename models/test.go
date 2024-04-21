package models

import (
	"time"
)

type Test struct {
	ID        int        `json:"id" gorm:"id"`
	Value     string     `json:"value" gorm:"value"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"deleted_at"`
}
