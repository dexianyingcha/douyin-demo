package main

import (
	"github.com/gin-gonic/gin"
	"simple-demo/model"
)

func main() {
	// 引用数据库
	model.Conn()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
