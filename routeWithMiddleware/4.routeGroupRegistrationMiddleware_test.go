package routeWithMiddleware

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func HeaderToken(msg string) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		if token == "1234" {
			context.Next()
			return
		}
		context.JSON(200, gin.H{"data": Response{0, msg, "token"}})
		context.Abort()
	}
}

func UserRouterToken(userGroup *gin.RouterGroup) {
	use := userGroup.Group("/token").Use(HeaderToken("token校验失败"))
	{
		use.GET("/user", UserListInfo)
	}
}

func TestRouteGroupRegistrationMiddleware(t *testing.T) {
	router := gin.Default()

	userGroup := router.Group("/group")

	UserRouterToken(userGroup)

	router.Run(":8080")
}
