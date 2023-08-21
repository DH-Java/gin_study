package view

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

type User struct {
	Name       string `json:"name" binding:"required,min=5"`
	Age        int    `json:"age"`
	Password   string `json:"password"`
	RePassword string `json:"re_password" binding:"eqfield=Password"`
}

func TestValidator(t *testing.T) {
	engine := gin.Default()
	var user User

	engine.POST("/validator", func(context *gin.Context) {
		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			log.Fatal("Validator", err)
		} else {
			context.JSON(http.StatusOK, user)
		}
	})

	engine.Run(":8080")
}
