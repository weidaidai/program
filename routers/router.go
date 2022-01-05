package routers

import (
	"program/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	controller := controller.Newcontrollers()
	s := r.Group("/student")
	{
		s.PUT("save", controller.Save)
		s.GET("select/:id", controller.Get)
		s.POST("update", controller.Update)
		s.DELETE("delete/:id", controller.Delete)
		s.GET("getall", controller.Getall)
	}
	return r
}
