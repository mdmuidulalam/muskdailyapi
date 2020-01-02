package main

import (
	"./routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	taskRoutes := (routes.Task{R: r.Group("/task")})
	taskRoutes.SetupRoutes()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
