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
	Saveredis(ctx *gin.Context)
	Studentrouter(r *gin.Engine)
}

//返回实现接口的对象
func New(svc database.StudentService, svc2 database.StudentService) Stucontroller {

	return &studentcontroller{svc_mysql: svc, svc_redis: svc2}

}

type studentcontroller struct {
	svc_mysql database.StudentService
	svc_redis database.StudentService
}

func (c *studentcontroller) Studentrouter(r *gin.Engine) {

	s := r.Group("/student")
	s.POST("/", c.Save)
	s.POST("/rdb", c.Saveredis)
	s.GET("/:id", c.Get)
	s.PUT("/", c.Update)
	s.DELETE("/:id", c.Delete)
	s.GET("/list", c.Getall)

}

//仅在redis保存
func (c *studentcontroller) Saveredis(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	var stu *model.Student
	if err := ctx.ShouldBindJSON(&stu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
	}
	err := c.svc_redis.SaveStudent(stu)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}

func (c *studentcontroller) Get(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	var std *model.Student
	std, err = c.svc_mysql.GetStudent(i)
	std, err = c.svc_redis.GetStudent(i)
	if std == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"ok": false, "data": nil})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": std})
}
func (c *studentcontroller) Getall(ctx *gin.Context) {
	var std []*model.Student
	std, err := c.svc_mysql.ListStudents()
	std, err2 := c.svc_redis.ListStudents()
	if std == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"ok": false, "data": nil})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err.Error()})
		return
	}
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": err2.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": std})

}
func (c *studentcontroller) Update(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	var stu *model.Student
	if err := ctx.ShouldBindJSON(&stu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
	}
	err := c.svc_mysql.UpdateStudent(stu)
	err = c.svc_redis.UpdateStudent(stu)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"ok": false, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}

func (c *studentcontroller) Save(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	var stu *model.Student
	if err := ctx.ShouldBindJSON(&stu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": err.Error()})
	}
	err := c.svc_mysql.SaveStudent(stu)
	// 同时在redis和mysql 保存
	//err2 := c.svc_redis.SaveStudent(stu)
	//if err2 != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"data": err2.Error()})
	//	return
	//}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}

func (c *studentcontroller) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	i, err2 := strconv.Atoi(id)
	if err2 != nil {
		return
	}
	err := c.svc_mysql.DeleteStudent(i)
	err2 = c.svc_redis.DeleteStudent(i)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}
