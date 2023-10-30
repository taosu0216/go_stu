package models

import "Book/utils"

func GetBookLists() []*BookInfo {
	lists := make([]*BookInfo, 300)
	utils.DB.Find(&lists)
	return lists
}
