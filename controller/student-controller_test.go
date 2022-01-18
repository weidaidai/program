package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStudentController_Get(t *testing.T) {

	c := &studentControllerImpl{}
	//切换到test模式
	gin.SetMode(gin.TestMode)
	//设置和注册路由
	r := gin.Default()
	r.GET("/mysql/list", c.GetAll)
	//测试的模拟请求
	req, err := http.NewRequest(http.MethodGet, "/mysql/list", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	// 创造响应路由
	w := httptest.NewRecorder()
	// 执行请求
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
