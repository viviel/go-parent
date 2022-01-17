package gin_test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

type Data struct {
	Id    int
	Value string
}

func Test1(t *testing.T) {
	r := gin.Default()
	setupRouter(r)
	err := r.Run()
	if err != nil {
		return
	}
}

func setupRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/post", func(c *gin.Context) {
		var j = Data{}
		err := c.BindJSON(&j)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
}
