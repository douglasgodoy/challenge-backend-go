package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	port := os.Getenv("PORT")
	router := gin.Default()
	initializeRoutes(router)
	router.Run(":" + port)
}
