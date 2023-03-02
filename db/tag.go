package db

import "gorm.io/gorm"

type Tag struct {
	Id string
	CategoryId int
	Name string
}


func CreateTag(tag Tag, db gorm.DB) {
	db.Create(tag)	
}

func GetTag(id string, db gorm.DB) (Tag, bool) {
	tag := Tag{}

	db.Where("id = ?", id).First(&tag)

	if tag.Id == "" {
		return tag, false
	}

	return tag, true
}
