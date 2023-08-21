package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

// 响应json
func _Json(context *gin.Context) {

	type userInfo struct {
		Name string `json:"name"`
		Age  string `json:"age"`
		//取消序列化
		Password string `json:"-"`
	}

	//Info := userInfo{"Mike", "20", "123456"}

	//context.JSON(http.StatusOK, Info)

	//usrs := map[string]string{"user_name": "Tom", "age": "30"}
	//返回map
	//context.JSON(http.StatusOK, usrs)

	//直接响应
	context.JSON(http.StatusOK, gin.H{"username": "Tom", "age": 33})

}

// 响应xml
func _xml(context *gin.Context) {
	context.XML(http.StatusOK, gin.H{"user": "xiaoMing", "status": http.StatusOK, "data": gin.H{"status": 299, "data": "nil"}})
}

// 响应yaml
func _yaml(context *gin.Context) {
	context.YAML(http.StatusOK, gin.H{"user": "xiaoWabf", "data": gin.H{"status": 300, "data": "nil"}})
}

// html模板
func _html(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{"user_name": "Mike"})
}

func _redirect(context *gin.Context) {
	//301永久重定向
	context.Redirect(301, "https://www.baidu.com")
	//302临时重定向
	context.Redirect(302, "https://xinghuo.xfyun.cn/desk")
}

func TestViewJson(t *testing.T) {
	router := gin.Default()

	//加载模板目录下所有的模板文件
	router.LoadHTMLGlob("template/*")

	//在golang中，没有相对文件的路径，他只有相对项目的路径
	//网页请求这个静态目录的前缀，第二个参数是一个目录，注意，前缀不要重复
	router.StaticFS("/static/document", http.Dir("static/document"))
	//配置单个文件，网页请求的路由，文件的路由
	router.StaticFile("/static/img", "static/img/*")

	router.GET("/json", _Json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/redirect", _redirect)

	router.Run(":8080")
}
