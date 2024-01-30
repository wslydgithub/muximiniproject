package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"miniproject/app/common/commonstruct"
	"miniproject/app/common/email"
	"miniproject/app/common/jwt1"
	"miniproject/app/model"
	"net/http"
)

// 获取验证码ok
func Getcode(r *gin.Engine) {
	r.POST("/saveplanets/signup/getcode", func(c *gin.Context) {
		commonstruct.Email = c.PostForm("username")
		email.Send(commonstruct.Email)
	})
}

// 检验+转到设置密码ok
func Solvehttpsignup(r *gin.Engine) {
	r.POST("/saveplanets/signup", func(c *gin.Context) {
		commonstruct.Codehttp = c.PostForm("testcode")
		if commonstruct.Codehttp == string(commonstruct.Code) {
			commonstruct.Usersign.Email = commonstruct.Email
			c.Redirect(http.StatusMovedPermanently, "saveplanets/signup/password")
		} else {
			c.String(http.StatusUnauthorized, "你的验证码有问题")
		}
	})
}

// 读取密码并注册+得到token ok
func Solvehttpsignpassword(r *gin.Engine) {
	r.POST("/saveplanets/signup/password", func(c *gin.Context) {
		commonstruct.Password = c.PostForm("password")
		commonstruct.Usersign.Password = commonstruct.Password
		var claims jwt1.Usercliams
		claims.Username = commonstruct.Email
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("muximiniproject"))
		if err != nil {
			fmt.Println("err(token):", err)
			return
		}
		c.Header("jwt1-token", tokenString)
		r.Use(jwt1.CORSMiddleware())
		c.Redirect(http.StatusMovedPermanently, "saveplanets/introduction")
	})
}

// 登录+得到token
func Solvelogin(r *gin.Engine) {
	r.POST("/saveplanets/login", func(c *gin.Context) {
		commonstruct.Email = c.PostForm("username")
		commonstruct.Password = c.PostForm("password")
		var user model.User
		result := commonstruct.DB.Where("name=?", commonstruct.Email).Where("password=?", commonstruct.Password).First(&user)
		if result.Error != nil {
			c.String(http.StatusUnauthorized, "你的邮箱或密码不正确")
			return
		} else {
			var claims jwt1.Usercliams
			claims.Username = commonstruct.Email
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte("muximiniproject"))
			if err != nil {
				fmt.Println("err(token):", err)
				return
			}
			c.Header("jwt1-token", tokenString)
			r.Use(jwt1.CORSMiddleware())
			c.Redirect(http.StatusMovedPermanently, "saveplanets/planet")
		}
	})
}

/*
	func Solveplanet(r *gin.Engine) {
		var r gin.Engine
		r.GET("Solveplanet")
	}
*/
func main() {
	r := gin.Default()
	Getcode(r)
	Solvehttpsignup(r)
	Solvehttpsignpassword(r)
	Solvelogin(r)
	r.Run(":8080")
}
