package dao
import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//初始化gorm
func InitDB() (err error) {
	dsn := "root:1234@(localhost:3306)/dy?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return nil
}
