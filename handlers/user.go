package handlers

import (
	"gin-weibo/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	data := models.GetUsers(page, size)

	meta := make(map[string]interface{})
	total, _ := models.GetUserTotal()
	meta["total"] = total
	meta["current_page"] = page
	meta["per_page"] = size
	meta["last_page"] = math.Ceil(float64(total/size)) + 1

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": data,
		"meta": meta,
	})
}
func UserCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "users/create.html", gin.H{
		"title": "注册",
	})
}
