package model

import "gorm.io/gorm"

func Paginate(startId, pageSize uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		return db.Where("id > ?", startId).Limit(int(pageSize))
	}
}
