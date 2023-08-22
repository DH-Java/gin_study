package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestDownLoadFile(t *testing.T) {

	engine := gin.Default()

	engine.GET("/download/:name", func(context *gin.Context) {
		fileName := context.Param("name")
		context.Header("Content-Type", "application/octet-stream")
		context.File(fmt.Sprintf("./uploadFile/%v", fileName))
		context.JSON(http.StatusOK, gin.H{"success": true})
	})

	engine.Run(":8080")
}
