package controller

import (
	"net/http"
	"program/database"
	"program/model"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Stucontroller interface {
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Save(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Getall(ctx *gin.Context)
	Studentrouter(r *gin.Engine)
}

//返回实现接口的对象
func Newmysql() Stucontroller {

	return &controller{}

}

type controller struct{ svc database.StudentService }

func (c *controller) Studentrouter(r *gin.Engine) {

	s := r.Group("/student")
	s.POST("/", c.Save)
	s.GET("/:id", c.Get)
	s.PUT("/", c.Update)
	s.DELETE("/:id", c.Delete)
	s.GET("/list_all", c.Getall)

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
		ctx.JSON(http.StatusNotFound, gin.H{"data": err.Error()})

	}
	ctx.JSON(http.StatusOK, gin.H{"data": u})

}
func (c *controller) Getall(ctx *gin.Context) {
	var std []*model.Student
	std, err := c.svc.ListStudents()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"data": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": std})

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
	}
	ctx.JSON(http.StatusOK, gin.H{"data": stu})
}

func (c *controller) Save(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	var stu *model.Student
	if err := ctx.ShouldBindJSON(&stu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
	err := c.svc.SaveStudent(stu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "save successfully"})
}

func (c *controller) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	i, err2 := strconv.Atoi(id)
	if err2 != nil {
		return
	}
	err := c.svc.DeleteStudent(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "delete successfully"})
}
