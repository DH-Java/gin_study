package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type User struct {
	Name       string   `json:"name" binding:"required,min=5"`
	Age        int      `json:"age"`
	Password   string   `json:"password"`
	RePassword string   `json:"re_password" binding:"eqfield=Password"`
	Sex        string   `json:"sex" binding:"oneof=man wumen,contains=a"`
	Str        string   `json:"str" binding:"contains=str"`
	StringList []string `json:"string_list" binding:"dive,startswith=str"`
	IP         string   `json:"ip" binding:"ip"`
	Url        string   `json:"url" binding:"url"`
	Uri        string   `json:"uri" binding:"uri"`
	Date       string   `json:"date" binding:"datetime=2006-01-02 15:04:06"`
}

func TestValidator(t *testing.T) {
	engine := gin.Default()
	var user User

	engine.POST("/validator", func(context *gin.Context) {
		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			//log.Fatal("Validator", err)
		} else {
			context.JSON(http.StatusOK, user)
		}
	})

	engine.Run(":8080")
}
