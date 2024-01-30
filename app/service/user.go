package service

import (
	"gorm.io/gorm"
	"miniproject/app/common/commonstruct"
	"miniproject/app/model"
)

func Adduser(db *gorm.DB, username string, planetname string, planetimage string) {
	var a model.User
	a.Useradd(commonstruct.Usersign)
	a.Planet[0].Initializeplanet(planetimage, planetname, username)
	db.Create(&a)
}
