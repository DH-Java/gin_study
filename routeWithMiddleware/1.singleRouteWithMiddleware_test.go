package routeWithMiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func m1(context *gin.Context) {
	fmt.Println("m1 ... start")
	context.Abort()
	context.Next()
	fmt.Println("m1 ... end")
}

func m2(context *gin.Context) {
	fmt.Println("m2 ... start")
	context.Next()
	fmt.Println("m2 ... end")
}

func m3(context *gin.Context) {
	fmt.Println("m3 ... start")
	context.Next()
	fmt.Println("m3 ... end")
}

func TestSingRouteWithMiddleware(t *testing.T) {
	engine := gin.Default()

	engine.GET("/route/middleware", m1, m2, m3)

	engine.Run(":8080")

}
