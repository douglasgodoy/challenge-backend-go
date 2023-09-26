package router

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	person(router)
	ping(router)
}

func ping(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func person(router *gin.Engine) {
	people := router.Group("people")
	{
		people.GET("/:id", handlers.GetPeople)
		people.POST("", handlers.CreatePerson)
	}
}
