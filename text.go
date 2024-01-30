package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User1 struct {
	MemberNumber string       `gorm:"primaryKey"`
	CreditCards  []CreditCard `gorm:"foreignKey:UserNumber"`
}

type CreditCard struct {
	Number     string
	UserNumber string `gorm:"index:size:255"`
}

func main() {
	dsn := "root:741074Hu050916@tcp(127.0.0.1:3306)/user"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//db.AutoMigrate(&User1{}, &CreditCard{})
	db.AutoMigrate(&User1{}, &CreditCard{})
	//var user User1
	//db.Preload("CreditCard").Find(&user)

}
