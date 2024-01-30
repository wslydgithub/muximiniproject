package model

type User struct {
	Name     string
	Password string
}

// 添加用户名
func (user *User) Useradd(user1 Usersl) {
	user.Name = user1.Email
	user.Password = user1.Password
}
