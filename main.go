package main

import (
	"miniproject/app/core/gorm"
	"miniproject/app/service"
)

func main() {
	var db = gorm.Linktodatabase()
	gorm.Migrate(db)
	service.Adduser(db, "lyd")
}
