package router

import (
	"gin-weibo/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	webRoute(r)
	apiRoute(r)
	return r
}
func webRoute(r *gin.Engine) *gin.Engine {
	//指定模板加载目录
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("views/**/*")
	r.GET("/", handlers.HomeIndex)
	users := r.Group("/users")
	{
		users.GET("/index", handlers.UserIndex)
	}
}
func apiRoute(r *gin.Engine) *gin.Engine {
	r.GET("/", handlers.HomeIndex)
	api := r.Group("/api")
	{
		api.GET("users/index", handlers.UserIndex)
	}
}
