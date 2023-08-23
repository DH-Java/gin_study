package routeWithMiddleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type User struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func matchToken(context *gin.Context) {
	var user User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, err.Error())
		return
	}
	if user.UserName != "admin" || user.Password != "1qazxsw2" {
		context.JSON(400, gin.H{"data": "用户校验失败"})
		context.Abort()
	}
	context.Set("user", User{
		user.UserName, user.Password,
	})
	context.Set("token", "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiI5MTUwMDEwN01BNjE1TFAzN1ciLCJjb250ZXh0VG9rZW5TZWNyZXQiOiJGQjRFMTRFMzE2OUZEQUUyQTNFRkEwOURFRjU2MjFDOCIsImNyZWF0ZVRpbWUiOjE2OTMzNzY4NjA3NTV9.KFi_oXFyBlnsFFUzrRGfKwG5pgq7fFlsmgpTjGg2IBU")
}

func loginSuccess(context *gin.Context) {
	value, _ := context.Get("token")
	_user, _ := context.Get("user")
	//断言
	user := _user.(User)
	context.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"token": value, "user": gin.H{"data": user}}})

}

func TestGlobalRegistrationMiddleware(t *testing.T) {
	router := gin.Default()

	router.Use(matchToken)

	router.POST("/global", loginSuccess)

	router.Run(":8080")
}
