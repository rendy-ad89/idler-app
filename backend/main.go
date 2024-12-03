package main

import (
	"idler/app/handlers"
	"idler/app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	// healthcheck API
	r := gin.Default()
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "healthy",
		})
	})

	// allow CORS
	r.Use(middleware.CORS())

	r.POST("/auth", handlers.ValidateAuth)
	r.POST("/users/register", handlers.CreateUser)

	// protected APIs
	api := r.Group("")
	api.Use(middleware.JWT())
	{
		api.GET("/machines", handlers.GetMachines)
		api.POST("/users/save-progress", handlers.SaveProgress)
		api.POST("/users/calc-profits", handlers.CalcOfflineProfits)
		api.GET("/users/machines", handlers.GetUsersMachines)
		api.POST("/users/machines", handlers.UpdateUsersMachines)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
