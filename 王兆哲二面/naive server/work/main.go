package main

import (
	"work/udb"
	"work/upage"

	"github.com/gin-gonic/gin"
)

func main() {

	//获取初始化的数据库
	db := udb.InitDB()
	//延迟关闭数据库
	defer db.Close()

	//yuanshen,启动！
	yuanshen := gin.Default()
	yuanshen.GET("/ping", upage.Ping)
	yuanshen.POST("/signup", upage.Signup)
	yuanshen.POST("/signin", upage.Signin)
	yuanshen.POST("/checkin", upage.Checkin)

	panic(yuanshen.Run(":8080"))
}
