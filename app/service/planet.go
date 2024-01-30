package service

import (
	"fmt"
	"gorm.io/gorm"
	"miniproject/app/model"
)

func Addplanet(db *gorm.DB, planetname string, username string, planetimage string) {
	var user model.User
	err := db.First(&user, "name=?", username).Error
	if err != nil {
		fmt.Println("can't find user:")
		return
	}
	user.Planetnumber++
	user.Planet[user.Planetnumber-1].Initializeplanet(planetimage, planetname, user.Name)
}

func Deleteplanet(db *gorm.DB, planetname string) {
	err1 := db.Exec("DELETE FROM animinals WHERE planetname =" + planetname).Error
	if err1 != nil {
		fmt.Println("err(delete animinal):", err1)
	}
	err2 := db.Exec("DELETE FROM plants WHERE planetname =" + planetname).Error
	if err2 != nil {
		fmt.Println("err(delete plant):", err1)
	}
	err3 := db.Exec("DELETE FROM goodbuildings WHERE planetname =" + planetname).Error
	if err3 != nil {
		fmt.Println("err(delete goodbuilding):", err1)
	}
	err4 := db.Exec("DELETE FROM badbuildings WHERE planetname =" + planetname).Error
	if err4 != nil {
		fmt.Println("err(delete badbuilding):", err1)
	}
	err5 := db.Exec("DELETE FROM reports WHERE planetname =" + planetname).Error
	if err5 != nil {
		fmt.Println("err(delete badbuilding):", err1)
	}
	err6 := db.Exec("DELETE FROM mainlands WHERE planetname =" + planetname).Error
	if err6 != nil {
		fmt.Println("err(delete badbuilding):", err1)
	}
	err7 := db.Exec("DELETE FROM planets WHERE name =" + planetname).Error
	if err7 != nil {
		fmt.Println("err(delete badbuilding):", err1)
	}

}
