package controller

import (
	"net/http"
	"program/database"
	"program/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

//controller 接口列表
type Studentcontroller interface {
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Save(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
type contorller struct {
	db database.MysqlStudentService
}

func Newstucontrollers() Studentcontroller {
	return &contorller{}
}
func (c *contorller) Get(ctx *gin.Context) {
	//var stu *model.Student
	id := ctx.Query("id")
	i, _ := strconv.Atoi(id)
	//获取查询参数
	num, err := c.db.GetStudent(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, num)
}

func (c *contorller) Update(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	var stu *model.Student
	if err := ctx.ShouldBindJSON(&stu); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	err := c.db.UpdateStudent(stu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "404"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "200"})
}

func (c *contorller) Save(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	var stu *model.Student
	if err := ctx.ShouldBindJSON(&stu); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	err := c.db.SaveStudent(stu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "404"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "200"})
}

func (c *contorller) Delete(ctx *gin.Context) {
	id := ctx.Query("id")
	i, _ := strconv.Atoi(id)
	err := c.db.DeleteStudent(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 200})
}
