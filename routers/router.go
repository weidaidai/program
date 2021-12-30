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
		s.PUT("/save")
		s.GET("/select", controller.Get)
		s.POST("/update")
		s.DELETE("/delete")
		s.GET("/getall")
	}
	return r
}
