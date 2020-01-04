package main

import (
	gin "github.com/gin-gonic/gin"
	config "muskdaily.com/config"
	routes "muskdaily.com/routes"
)

func main() {
	configuration := config.GetConfiguration()
	r := gin.Default()
	taskRoutes := (routes.Task{R: r.Group("/task")})
	taskRoutes.SetupRoutes()
	r.Run(":" + configuration.Port)
}
