package models

import (
	"Book/utils"
	"gorm.io/gorm"
)

func GetBookLists() []*BookInfo {
	lists := make([]*BookInfo, 300)
	utils.DB.Find(&lists)
	return lists
}
func FindBookByID(id uint) (bool, *BookInfo, *gorm.DB) {
	book := &BookInfo{}
	data := utils.DB.Where("id = ?", id).First(book)
	if book.ID != 0 {
		return true, book, data
	}
	return false, book, data
}
