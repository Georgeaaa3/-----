package upage

import (
	"net/http"
	"work/udb"
	"work/utoken"

	"github.com/gin-gonic/gin"
)

func Checkin(ctx *gin.Context) {
	db := udb.GetDB()

	//输入
	access_token := ctx.PostForm("access_token")
	checkword := ctx.PostForm("checkword")

	if len(access_token) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 311,
			"msg":  "请输入token",
			"data": access_token,
		})
		return
	}

	//解码
	c, err := utoken.ParseToken(access_token)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 312,
			"msg":  "无效token",
			"data": access_token,
		})
		return
	}

	var user udb.User
	db.Where("Username = ?", c.Uname).First(&user)
	user.Point += len(checkword)

	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		"name":  c.Uname,
		"point": user.Point,
	})
	db.Save(&user)
}
