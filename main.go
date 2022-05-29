package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/gin-gonic/gin"
)

var U1 controller.User
var U2 controller.User

func main() {
	r := gin.Default()
	initRouter(r)
	dao.InitDB()

	dao.DB.AutoMigrate(&controller.User{})
	dao.DB.AutoMigrate(&controller.Comment{})
	dao.DB.AutoMigrate(&controller.Video{})
	dao.DB.AutoMigrate(&controller.ClickRelation{})
	dao.DB.AutoMigrate(&controller.Relation{})

	controller.InitDemo() //初始化测试数据

	defer dao.DB.Close()
	r.Run()

}
