package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//指定模板加载目录

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("views/**/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "首页",
		})
	})
	router.Run()
}
