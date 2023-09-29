package upage

import (
	"net/http"
	"work/udb"
	"work/utoken"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func checkpassward(passward string) int {
	a, b, c := 0, 0, 0
	for _, ch := range passward {
		if ch >= 'a' && ch <= 'z' {
			a = 1
		}
		if ch >= 'A' && ch <= 'Z' {
			b = 1
		}
		if ch >= '0' && ch <= '9' {
			c = 1
		}
	}
	return a + b + c
}

func Signup(ctx *gin.Context) {

	db := udb.GetDB()

	//输入
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//处理用户
	if len(username) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 111,
			"msg":  "用户名不能为空",
			"data": username,
		})
		return
	}

	var user udb.User
	db.Where("Username = ?", username).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 112,
			"msg":  "用户已存在",
			"data": username,
		})
		return
	}

	//处理密码
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 113,
			"msg":  "密码不能少于6位",
			"data": password,
		})
		return
	}

	if checkpassward(password) < 3 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 114,
			"msg":  "密码需包括大小写字母和数字",
			"data": password,
		})
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 101,
			"msg":  "密码加密错误",
			"data": password,
		})
		return
	}

	//添加用户
	newUser := udb.User{
		Username: username,
		Password: string(hasedPassword),
		Point:    0,
	}
	db.Create(&newUser)

	toke, err := utoken.GetnerateToken(newUser.Username)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"access_token": toke,
		})
	}
}
