package models

import (
	"gorm.io/gorm"
)

type BookInfo struct {
	gorm.Model
	Name         string
	Author       string
	Img          string
	Tags         string
	Description  string
	BorrowerName string
	BorrowedTime string
	ReturnTime   string
	IsReturn     string
}

func (table *BookInfo) TableName() string {
	return "book_infos"
}
