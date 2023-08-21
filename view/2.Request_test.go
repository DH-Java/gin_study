package view

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"testing"
)

func _query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	if v, ok := c.GetQuery("user"); ok {
		fmt.Println(v)
	} else {
		log.Print("GetQuery(未有user参数)")
	}
	//拿到多个相同的查询参数
	fmt.Println(c.QueryArray("user"))
}

func _params(c *gin.Context) {
	//动态参数
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	if v, err := c.GetPostFormArray("name"); err {
		fmt.Println(v)
	} else {
		fmt.Println("没有传参")
	}
	fmt.Println(c.DefaultPostForm("addr", "重庆市"))
	fmt.Println(c.MultipartForm())
}

// 原始参数
func _raw(c *gin.Context) {
	var user struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	if err := bindJson(c, &user); err != nil {
		log.Fatal("bindJson", err)
	}
	fmt.Println(user)
}

func bindJson(c *gin.Context, obj any) error {
	if data, err := c.GetRawData(); err != nil {
		fmt.Println("c,GetRawData:", err)
	} else {
		contenType := c.GetHeader("Content-Type")
		switch contenType {
		case "application/json":
			if err := json.Unmarshal(data, &obj); err != nil {
				log.Fatal("json Unmarshal", err)
				return err
			}
		}
	}
	return nil
}

func TestRequest(t *testing.T) {
	engine := gin.Default()

	engine.GET("/query", _query)
	engine.GET("/param/:user_id", _params)
	engine.GET("/param/:user_id/:book_id", _params)
	engine.GET("/form", _form)
	engine.GET("/raw", _raw)

	engine.Run(":8080")
}
