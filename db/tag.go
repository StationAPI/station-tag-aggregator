package db

import "gorm.io/gorm"

type Tag struct {
	Id         string 
	CategoryId string `json:"category_id"`
	Name       string `json:"name"`
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

func ListTags(db gorm.DB) []Tag {
	var tags []Tag

	db.Find(&tags)	

	return tags
}
