package Models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

func (b *Book) TableName() string {
	return "book"
}
