package view

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"testing"
)

type UserInfos struct {
	Name string `json:"name" binding:"required,sign" msg:"用户名授权失败"`
	Age  string `json:"age" binding:"required" msg:"密码授权失败"`
}

// GetValidMsg 返回结构体中的msg参数
func GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, fieldError := range errs {
			if f, exits := getObj.Elem().FieldByName(fieldError.Field()); exits {
				return f.Tag.Get("msg")
			}
		}
	}
	return err.Error()
}

func TestException(t *testing.T) {
	engine := gin.Default()

	engine.GET("/exception", func(context *gin.Context) {
		var user UserInfos
		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusInternalServerError, GetValidMsg(err, &user))
			return
		} else {
			context.JSON(http.StatusOK, user)
		}
	})

	engine.Run(":8080")
}
