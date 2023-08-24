package document

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestAsciiJson(t *testing.T) {
	router := gin.Default()

	router.GET("/asciiJson", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "<br>",
		}
		//使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
		//{"lang":"Go\u8bed\u8a00","tag":"\u003cbr\u003e"}
		context.AsciiJSON(http.StatusOK, data)
	})
	router.Run()
}
