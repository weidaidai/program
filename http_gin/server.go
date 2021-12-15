package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"program/http_gin/controllers"
)

func main() {
	//创路由
	r := gin.Default()
	//get 方法 ，调用回调函数
	r.GET("/ping", func(cxt *gin.Context) {
		cxt.String(200, "%s", "pong")
	})
	//获取文件,在访问参数要后要/video/具体文件名
	r.Static("/video", "./video")
	//具体文件路径
	videocontroller := controllers.Newvidecontrollers()

	//路由组
	videoGroup := r.Group("/videos")
	//videoGroup.Use(middlewares.Logger())

	videoGroup.GET("/", videocontroller.Get)

	videoGroup.PUT("/:id", videocontroller.Update)

	videoGroup.DELETE("/:id", videocontroller.Delete)

	videoGroup.POST("/", videocontroller.Create)

	videoGroup.StaticFile("/myvideo", "./video/Road.mp4")
	log.Fatalln(r.Run("127.0.0.1:8080"))

}
