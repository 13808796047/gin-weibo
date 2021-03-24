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
	r.GET("users", handlers.UserIndex)
	r.GET("signup", handlers.UserCreate)
	return r
}
func apiRoute(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	{
		api.GET("users", handlers.UserIndex)
	}
	return r
}
