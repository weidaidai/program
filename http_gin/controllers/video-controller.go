package controllers

import (
	"github.com/gin-gonic/gin"
	"program/http_gin/models"
	"sync"
)

//controller 接口列表
type videocontroller interface {
	Get(ctx *gin.Context)
	Update(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type videoid struct {
	counter int
	mtx     sync.Mutex
}

//

func (v *videoid) getnewid() int {
	//加锁（wed server 并发执行 避免重复需要加锁）
	v.mtx.Lock()
	defer v.mtx.Unlock()
	v.counter++
	return v.counter

}

var v *videoid = &videoid{}

//
type controller struct {
	video []models.Video //类似数据库
}

func Newvidecontrollers() videocontroller {
	return &controller{video: make([]models.Video, 0)}
}

//实现方法
func (c *controller) Get(ctx *gin.Context) {
	ctx.JSON(200, c.video)
}

func (c *controller) Update(ctx *gin.Context) {
	//从controller拿然后更新
	var videoupdate models.Video

	if err := ctx.ShouldBindUri(&videoupdate); err != nil {
		ctx.String(400, "请求错误%v", err)
		return
	}
	for id, video := range c.video {
		if video.Id == videoupdate.Id {

			c.video[id] = videoupdate
			ctx.String(200, "已成功更新%d",videoupdate.Id)
			return
		}
	}
	ctx.String(400, "bad request cannot find video with %d to update", videoupdate.Id)
}
func (c *controller) Create(ctx *gin.Context) {
	video := models.Video{Id: v.getnewid()}
	if err := ctx.BindJSON(&video); err != nil {
		ctx.String(400, "请求错误%v", err)

	}
	c.video = append(c.video, video)
	ctx.String(200, "成功创建新的视频%d", video.Id)

}
func (c *controller) Delete(ctx *gin.Context) {
	var videodelete models.Video

	if err := ctx.ShouldBindUri(&videodelete); err != nil {
		ctx.String(400, "请求错误%v", err)
		return
	}
	for id, video := range c.video {
		if video.Id == videodelete.Id {

			c.video = append(c.video[0:id], c.video[id+1:len(c.video)]...)
			ctx.String(200, "已成功删除%d",videodelete.Id)
			return
		}
	}
	ctx.String(400, "bad request cannot find video with %d to delete", videodelete.Id)
}
