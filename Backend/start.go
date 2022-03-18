package main

import (
	"URL-Sort/Backend/auth"
	"URL-Sort/Backend/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(auth.CORSMiddleware())
	router.GET("/:url_id", controllers.GetRedirect)
	router.POST("/api/v1/urls", controllers.SetURL)
	router.Run(":8687")
}
