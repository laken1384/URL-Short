package main

import (
	"Backend/auth"
	"Backend/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(auth.CORSMiddleware())
	router.GET("/api/:url_id", controllers.GetRedirect)
	router.POST("/api/v1/urls", controllers.SetURL)
	router.Run(":8687")
}
