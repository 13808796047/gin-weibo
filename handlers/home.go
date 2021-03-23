package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"title": "首页",
	})
}
