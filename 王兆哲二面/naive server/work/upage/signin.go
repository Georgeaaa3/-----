package upage

import (
	"net/http"
	"work/udb"
	"work/utoken"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signin(ctx *gin.Context) {

	db := udb.GetDB()

	//输入
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//处理用户
	var user udb.User
	db.Where("Username = ?", username).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 211,
			"msg":  "用户不存在",
			"data": username,
		})
		return
	}

	//处理密码
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    212,
			"message": "密码不能少于6位",
			"data":    password,
		})
		return
	}

	if checkpassward(password) < 3 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 213,
			"msg":  "密码需包括大小写字母和数字",
			"data": password,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 214,
			"msg":  "密码错误",
			"data": password,
		})
		return
	}

	toke, err := utoken.GetnerateToken(user.Username)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"access_token": toke,
		})
	}
}
