package log

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	//输出日志到文件
	logFile, _ := os.Create("gin.log")

	//如果需要同时将日志
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	router := gin.Default()

	router.GET("logInfo", func(context *gin.Context) {

	})

	router.Run(":8080")
}
