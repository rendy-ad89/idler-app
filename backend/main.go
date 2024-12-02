package main

import (
	"idler/app/handlers"
	"idler/app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	r.Use(middleware.CORS())

	r.POST("/auth", handlers.ValidateAuth)
	r.POST("/users/register", handlers.CreateUser)

	api := r.Group("")
	api.Use(middleware.JWT())
	{
		api.GET("/machines", handlers.GetMachines)
		api.POST("/users/save-progress", handlers.SaveProgress)
		api.POST("/users/calc-profits", handlers.CalcOfflineProfits)
		api.POST("/users/machines", handlers.GetUsersMachines)
		api.POST("/users/machines/update", handlers.UpdateUsersMachines)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
