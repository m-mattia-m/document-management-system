package models

import (
	"gorm.io/gorm"
	"time"
)

type Docuemnt struct {
	gorm.Model `swaggerignore:"true"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	Date       time.Time `json:"date"`
}
