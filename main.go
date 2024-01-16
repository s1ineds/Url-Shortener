package main

import (
	"url-shortener/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("views/*")

	router.GET("/index", controllers.IndexController)
	router.POST("/index", controllers.PostIndexForm)
	router.GET("/:urlId", controllers.RedirectIndexControlller)

	router.GET("/api/getUrls", controllers.GetUrls)
	router.GET("/api/getShortUrl/:originalUrl", controllers.GetShortUrl)

	router.Run("127.0.0.1:80")
}
