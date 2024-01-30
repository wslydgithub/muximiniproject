package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"miniproject/app/model"
)

func Linktodatabase() *gorm.DB {
	dsn := "root:741074Hu050916@tcp(127.0.0.1:3306)/user"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Planet{})
	db.AutoMigrate(&model.Planet{}, &model.Mainland{})
	db.AutoMigrate(&model.Mainland{}, &model.Plant{})
	db.AutoMigrate(&model.Mainland{}, &model.Goodbuilding{})
	db.AutoMigrate(&model.Mainland{}, &model.Badbuilding{})
	db.AutoMigrate(&model.Mainland{}, &model.Animinal{})
	db.AutoMigrate(&model.Mainland{}, &model.Report{})
	db.AutoMigrate(&model.Animinal{})
	db.AutoMigrate(&model.Plant{})
	db.AutoMigrate(&model.Goodbuilding{})
	db.AutoMigrate(&model.Badbuilding{})
	db.AutoMigrate(&model.Report{})
}
