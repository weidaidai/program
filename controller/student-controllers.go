package controller

import (
	"net/http"
	"program/database"
	"program/model"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type StudentController interface {
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Save(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	StudentRouter(r *gin.Engine, prefix string)
}

//返回实现接口的对象
func New(svc database.StudentService) StudentController {
	return &studentControllerImpl{svc: svc}
}

type studentControllerImpl struct {
	svc database.StudentService
}

func (c *studentControllerImpl) StudentRouter(r *gin.Engine, prefix string) {

	s := r.Group("/student" + prefix)
	s.POST("/", c.Save)
	s.GET("/:id", c.Get)
	s.PUT("/:id", c.Update)
	s.DELETE("/:id", c.Delete)
	s.GET("/list", c.GetAll)

}

func (c *studentControllerImpl) Get(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	switch i {

	}
	var std *model.Student
	std, err = c.svc.GetStudent(i)
	if std == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"data": std})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": true, "data": std})
}
func (c *studentControllerImpl) GetAll(ctx *gin.Context) {
	var stds []*model.Student
	stds, err := c.svc.ListStudents()
	if stds == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"data": stds})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": true, "data": stds})

}
func (c *studentControllerImpl) Update(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	id := ctx.Params.ByName("id")
	i, err2 := strconv.Atoi(id)
	if err2 != nil {
		return
	}
	stu := &model.Student{}
	stu.Id = i
	if err := ctx.ShouldBindJSON(stu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
	}

	err := c.svc.UpdateStudent(stu)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"ok": false, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}

func (c *studentControllerImpl) Save(ctx *gin.Context) {
	// 获取传递的参数 转换成 struct
	stu := &model.Student{}
	if err := ctx.ShouldBindJSON(&stu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": err.Error()})
	}
	err := c.svc.SaveStudent(stu)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}

func (c *studentControllerImpl) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	i, err2 := strconv.Atoi(id)
	if err2 != nil {
		return
	}

	err := c.svc.DeleteStudent(i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true})
}
