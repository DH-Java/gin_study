package view

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

type Books struct {
	BookName string `json:"bookName"`
	Price    string `json:"price"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func BindJson(c *gin.Context, obj any) error {
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

func _getList(c *gin.Context) {
	books := []Books{
		{"Go语言入门", "88"},
		{"Python语言入门", "77"},
		{"Java语言入门", "99"},
	}
	c.JSON(http.StatusOK, Response{http.StatusOK, books, "成功"})
}
func _getDetail(c *gin.Context) {
	fmt.Println(c.Param("id"))

	c.JSON(http.StatusOK, Response{http.StatusOK, Books{
		"钢铁是怎样炼成的",
		"50",
	}, "成功"})

}
func _getCreate(c *gin.Context) {
	var book Books
	if err := BindJson(c, &book); err != nil {
		log.Fatal("BindJson Error:", err)
	} else {
		c.JSON(http.StatusOK, Response{http.StatusOK, book, "添加成功"})
	}
}
func _getUpdate(c *gin.Context) {
	fmt.Println(c.Param("id"))
	var book Books
	if err := BindJson(c, &book); err != nil {
		log.Fatal("BindJson Error:", err)
	} else {
		c.JSON(http.StatusOK, Response{http.StatusOK, book, "修改成功"})
	}
}
func _getDelete(c *gin.Context) {
	fmt.Println(c.Param("id"))
	c.JSON(http.StatusOK, Response{http.StatusOK, 1, "删除成功"})

}

func TestRestFul(t *testing.T) {
	engine := gin.Default()

	engine.GET("/articles", _getList)
	engine.GET("/articles/:id", _getDetail)
	engine.POST("/articles", _getCreate)
	engine.PUT("/articles/:id", _getUpdate)
	engine.DELETE("/articles/:id", _getDelete)

	engine.Run(":8080")
}
