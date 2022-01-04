package routers

import (
	"program/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Router() *gin.Engine {
	r := gin.Default()

	s := r.Group("/student")
	{
		s.PUT("save", controller.Save)
		s.GET("select:id", controller.Get)
		s.POST("update", controller.Update)
		s.DELETE("delete:id", controller.Delete)
	}
	return r
}
