package view

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"testing"
)

// 自定义验证器
func signValid(fl validator.FieldLevel) bool {
	names := []string{"Mike", "Tom"}
	for _, v := range names {
		name := fl.Field().Interface().(string)
		if name == v {
			return false
		}
	}
	return true
}

func TestCustomValidator(t *testing.T) {
	engine := gin.Default()

	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate.RegisterValidation("sign", signValid)
	}

	engine.GET("/customValidator", func(context *gin.Context) {
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
