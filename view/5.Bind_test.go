package view

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func TestBind(t *testing.T) {
	engine := gin.Default()

	engine.POST("/shouldBindJSON", func(context *gin.Context) {
		var userInfo UserInfo
		if err := context.ShouldBindJSON(&userInfo); err != nil {
			log.Fatal("ShouldBindJSON", err)
		} else {
			context.JSON(http.StatusOK, userInfo)
		}
	})

	engine.POST("/shouldBindQuery", func(context *gin.Context) {
		var userInfo UserInfo
		if err := context.ShouldBindQuery(&userInfo); err != nil {
			log.Fatal("ShouldBindQuery", err)
		} else {
			context.JSON(http.StatusOK, userInfo)
		}
	})

	engine.POST("/shouldBindUri/:name/:age/:sex", func(context *gin.Context) {
		var userInfo UserInfo
		if err := context.ShouldBindUri(&userInfo); err != nil {
			log.Fatal("ShouldBindUri", err)
		} else {
			context.JSON(http.StatusOK, userInfo)
		}
	})

	//绑定form-data
	engine.POST("/shouldBind/form", func(context *gin.Context) {
		var userInfo UserInfo
		if err := context.ShouldBind(&userInfo); err != nil {
			log.Fatal("ShouldBind ", err)
		} else {
			context.JSON(http.StatusOK, userInfo)
		}
	})

	engine.Run(":8080")

}
