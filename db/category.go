package db

import "gorm.io/gorm"

type Category struct {
	Id   string
	Name string
	Tags []Tag
}

func CreateCategory(category Category, db gorm.DB) {
	db.Create(category)
}

func GetCategory(id string, db gorm.DB) (Category, bool) {
	category := Category{}

	db.Where("id = ?", id).First(&category)

	if category.Id == "" {
		return category, false
	}

	return category, true
}
