package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}

func main() {
	//创建一给默认的路由
	router := gin.Default()
	//绑定路由规则，访问/index的路由，将有对应的函数去处理
	router.GET("/index", Index)

	//启动监听，gin会把web服务运行在本机的0.0.0.0:8080端口上
	router.Run(":8080")

	//用原生http服务的方式,router.Run本质就是http.ListenAndServe的进一步封装
	http.ListenAndServe(":8081", router)
}
