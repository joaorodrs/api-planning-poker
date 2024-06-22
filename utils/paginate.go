package utils

import (
	"github.com/jinzhu/gorm"
)

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var totalCount int
		db.Model(value).Count(&totalCount)
		pagination.Total = totalCount

		offset := (pagination.Page - 1) * pagination.Take
		db = db.Offset(offset).Limit(pagination.Take)

		if pagination.Take > 0 {
			pagination.LastPage = (totalCount + pagination.Take - 1) / pagination.Take
		} else {
			pagination.LastPage = 0
		}

		return db
	}
}

type Pagination struct {
	Take     int `json:"take"`
	Page     int `json:"page"`
	Total    int `json:"total"`
	LastPage int `json:"last_page"`
}
