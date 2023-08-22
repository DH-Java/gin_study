package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"testing"
)

func SaveUploadFile(context *gin.Context) {
	if file, err := context.FormFile("file"); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		context.SaveUploadedFile(file, fmt.Sprintf("uploadFile/%v", file.Filename))
		context.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"fileName": file.Filename, "fileSize": file.Size / 1024}})
	}
}

// 读取上传文件
func ReadFile(context *gin.Context) {
	if file, err := context.FormFile("file"); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		if opne, err := file.Open(); err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		} else {
			reader, _ := io.ReadAll(opne)
			context.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"fileName": file.Filename, "fileSize": file.Size / 1024, "data": string(reader)}})
		}

	}
}

// 创建拷贝上传
func CreateCopy(context *gin.Context) {
	if file, err := context.FormFile("file"); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		if opne, err := file.Open(); err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
			return
		} else {
			if create, err := os.Create(fmt.Sprintf("./uploadFile/%v", file.Filename)); err != nil {
				context.JSON(http.StatusInternalServerError, err.Error())
				return
			} else {
				defer create.Close()
				if written, err := io.Copy(create, opne); err != nil {
					context.JSON(http.StatusInternalServerError, err.Error())
					return
				} else {
					context.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"fileName": file.Filename, "fileSize": file.Size / 1024, "data": written}})
				}
			}
		}
	}
}

// 批量上传
func Uploads(context *gin.Context) {
	if file, err := context.MultipartForm(); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		opne := file.File["uploads"]
		for _, value := range opne {
			context.SaveUploadedFile(value, fmt.Sprintf("./uploadFile/%v", value.Filename))
		}
		context.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"size": len(opne)}})
	}
}

func TestUpload(t *testing.T) {
	engine := gin.Default()

	engine.POST("/uploadFile", SaveUploadFile)
	engine.POST("/readFile", ReadFile)
	engine.POST("/createCopy", CreateCopy)
	engine.POST("/uploads", Uploads)

	engine.Run(":8080")

}
