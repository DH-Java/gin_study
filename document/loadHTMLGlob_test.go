package document

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestLoadHTMLGlob(t *testing.T) {
	router := gin.Default()
	router.LoadHTMLGlob("../templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.Run(":8080")
}
