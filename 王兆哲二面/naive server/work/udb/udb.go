package udb

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Point    int
}

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "user_db" //需要先用mysql建立数据库
	username := "root"    //自己的名称
	password := "225639"  //自己的密码
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true", username, password, host, port, database, charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	db.AutoMigrate(&User{})
	DB = db
	return db
}
