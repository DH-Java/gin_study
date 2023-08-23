package routeWithMiddleware

import (
	"github.com/gin-gonic/gin"
	"testing"
)

type UserInfo struct {
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func UserListInfo(context *gin.Context) {
	userList := []UserInfo{
		{"Mike", 20}, {"Tom", 21},
	}
	context.JSON(200, gin.H{"data": Response{0, userList, "userInfo"}})
}

func UserRouter(userGroup *gin.RouterGroup) {
	userGroup.GET("/group", UserListInfo)
}

func TestRouteGroup(t *testing.T) {
	router := gin.Default()

	userGroup := router.Group("/user")

	UserRouter(userGroup)

	router.Run(":8080")
}
