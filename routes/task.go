package routes

import (
	"github.com/gin-gonic/gin"
)

type Task struct {
	R *gin.RouterGroup
}

func (this *Task) SetupRoutes() {
	this.R.GET("get", this.GetTask)
}

func (this *Task) GetTask(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Success",
	})
}
