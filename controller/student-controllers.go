package controller

import (
	"net/http"
	"program/database"
	"program/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	i, _ := strconv.Atoi(id)
	db := database.MysqlStudentService{}
	u, err := db.GetStudent(i)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {

		ctx.JSON(http.StatusOK, u)
	}
}

func Update(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	c := database.MysqlStudentService{}
	var stu *model.Student
	if err := ctx.ShouldBindJSON(&stu); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	err := c.UpdateStudent(stu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "404"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "200"})
}

func Save(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	c := database.MysqlStudentService{}
	var stu *model.Student
	if err := ctx.ShouldBindJSON(&stu); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	err := c.SaveStudent(stu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "404"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "200"})
}

func Delete(ctx *gin.Context) {
	c := database.MysqlStudentService{}
	id := ctx.Query("id")
	i, _ := strconv.Atoi(id)
	err := c.DeleteStudent(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 200})
}
