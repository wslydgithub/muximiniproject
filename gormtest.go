package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//小小的总结一下，注意几点
//1.注意大小写在代码中的都要开头大写，而sql中的原样写
//2.注意&User{}
/*type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:741074Hu050916@tcp(127.0.0.1:3306)/user"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}*/

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	//连接mysql数据库
	dsn := "root:741074Hu050916@tcp(127.0.0.1:3306)/user"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//建一个表格
	db.AutoMigrate(&User{})
	//增加
	/*user1 := []User{
		User{Name: "覃玉荣", Age: 18},
		User{Name: "梁雅栋", Age: 18},
	}
	db.Create(user1)

	//查找
	var user_text User
	db.First(&user_text, "name=", "覃玉荣")
	//改变
	db.Model(&User{}).Where("age=?", 19).Update("age", 17)
	//伪装删除
	db.Where("name=?", "梁雅栋").Delete(&User{})
	//回归
	var users []User
	db.Clauses(clause.Returning{}).Where("name = ?", "覃玉荣").Delete(&users)
	//永久删除
	db.Unscoped().Where("name=?", "NULL").Delete(&User{})*/
}
