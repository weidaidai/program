package routers

import (
	"program/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	controller := controller.Newstucontrollers()
	s := r.Group("/student")
	{
		s.PUT("/save", controller.Save)
		s.GET("/select", controller.Get)
		s.POST("/update", controller.Update)
		s.DELETE("/delete", controller.Delete)
	}
	return r
}
