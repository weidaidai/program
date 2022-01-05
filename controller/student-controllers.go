package controller

import (
	"fmt"
	"net/http"
	"program/database"
	"program/model"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type stucontroller interface {
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Save(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Getall(ctx *gin.Context)
}

func Newcontrollers() stucontroller {
	return &controller{}
}

type controller struct {
	svc database.MysqlStudentService
}

func (c *controller) Get(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	var u *model.Student
	u, err = c.svc.GetStudent(i)
	if err != nil {
		ctx.JSON(http.StatusNotFound, u)
	} else {
		ctx.JSON(http.StatusOK, u)
	}
}
func (c *controller) Getall(ctx *gin.Context) {
	var std []*model.Student
	std, err := c.svc.ListStudents()
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, std)
	}
}
func (c *controller) Update(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	var stu *model.Student
	err := c.svc.UpdateStudent(stu)
	if err := ctx.ShouldBindJSON(&stu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Println(stu)
		ctx.JSON(http.StatusOK, stu)
	}
}

func (c *controller) Save(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	var stu *model.Student
	if err := ctx.ShouldBindJSON(&stu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	err := c.svc.SaveStudent(stu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "404"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "200"})
}

func (c *controller) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	i, err2 := strconv.Atoi(id)
	if err2 != nil {
		return
	}
	err := c.svc.DeleteStudent(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "del success"})
}
