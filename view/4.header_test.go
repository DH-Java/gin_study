package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"testing"
)

func TestHeader(t *testing.T) {
	engine := gin.Default()

	engine.GET("/header", func(context *gin.Context) {
		//获取Header中的值，不区分大小写，但只获取第一个
		fmt.Println(context.GetHeader("User-Agent"))
		fmt.Println(context.GetHeader("xxx"))

		//获取所有Header用map[string] []string返回
		fmt.Println(context.Request.Header)

		//如果使用Get()方法，不区分大小写，也是只获取第一个value
		fmt.Println(context.Request.Header.Get("xxx"))

		//map[Key]需要区分大小写，但是可以获取多个
		fmt.Println(context.Request.Header["Xxx"])

		context.JSON(200, "获取header的值")
	})

	engine.GET("/python", func(context *gin.Context) {
		if strings.Contains(context.GetHeader("python"), "python") {
			context.JSON(http.StatusOK, "python数据")
		} else {
			context.JSON(http.StatusOK, "用户数据")
		}
	})

	engine.GET("/setHeader", func(context *gin.Context) {
		context.Header("token", "3211dwe1231")
		context.Header("Content-Type", "application/text; charset=utf-8")
		context.JSON(http.StatusOK, "设置响应头")
	})

	engine.Run(":8080")
}
