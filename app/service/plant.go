package service

import (
	"gorm.io/gorm"
	"miniproject/app/model"
)

func addanimal(db *gorm.DB) {
	db.Create(&model.Animinal{Name: ""})
}
