package main

import (
	"gin-weibo/models"
	"gin-weibo/router"
)

func init() {
	models.Setup()
}
func main() {
	router := router.SetupRouter()

	router.Run()
}
