package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var err error

// 连接数据库
func Conn() {
	//db, err := gorm.Open("mysql", "root:wtt991221@tcp(192.168.1.235:3306)/douyin") //实验室的ip
	//db, err := gorm.Open("mysql", "root:wtt991221@tcp(172.26.84.206:3306)/douyin") //宿舍的ip
	//db, err := gorm.Open("mysql", "root:wtt991221@tcp(172.19.37.110:0)/douyin") //教学楼的ip
	db, err = gorm.Open("mysql", "root:whh666?@tcp(127.0.0.1:3306)/douyin?parseTime=true") //本地
	if err != nil {
		fmt.Print("sql open error", err)
		//return nil
	}

	db.SingularTable(true)
	db.LogMode(true)
	//return db
}
