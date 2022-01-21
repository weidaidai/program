package controller

import (
	"net/http"
	"program/model"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

var testUrl string = "http://127.0.0.1:8080"

func TestStudentController_redisPostok(t *testing.T) {
	e := httpexpect.New(t, testUrl) //创建一个httpexpect实例
	//stu := &model.Student{Id: 18, Name: "555", Age: 18}
	postData := map[string]interface{}{ //创建一个json变量
		"id":   18,
		"name": "555",
		"age":  18,
	}
	contentType := "application/json;charset=utf-8"

	e.POST("/student/redis/"). //post 请求
					WithHeader("ContentType", contentType). //定义头信息
					WithJSON(postData).                     //传入json body
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().              //json body实例化
					ContainsKey("ok").     //检验是否包括key
					ValueEqual("ok", true) //对比key的value

}

func TestStudentController_redisGetOk(t *testing.T) {
	stu := &model.Student{Id: 18, Name: "555", Age: 18}
	e := httpexpect.New(t, testUrl) //创建一个httpExpect实例
	e.GET("/student/redis/18").     //get请求
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().               //json body实例化
					ContainsKey("data").    //检验是否包括key
					ValueEqual("data", stu) //对比key的value

}
func TestStudentController_redisGetFail(t *testing.T) {
	e := httpexpect.New(t, testUrl) //创建一个httpExpect实例
	e.GET("/student/mysql/88").     //ge请求
					Expect().
					Status(http.StatusNotFound). //判断请求是否404
					JSON().
					Object().                  //json body实例化
					ContainsKey("data").       //检验是否包括key
					ValueEqual("data", "null") //对比key的value
}

func TestStudentController_redisPutok(t *testing.T) {

	e := httpexpect.New(t, testUrl) //创建一个httpexpect实例
	//stu := &model.Student{Id: 18, Name: "555", Age: 18}
	postData := map[string]interface{}{ //创建一个json变量
		"name": "xiao",
		"age":  88,
	}
	contentType := "application/json;charset=utf-8"

	e.PUT("/student/redis/18"). //put 请求
					WithHeader("ContentType", contentType). //定义头信息
					WithJSON(postData).                     //传入json body
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().              //json body实例化
					ContainsKey("ok").     //检验是否包括key
					ValueEqual("ok", true) //对比key的value

}

func TestStudentController_redisGetList(t *testing.T) {
	want := []*model.Student{{Id: 18, Name: "xiao", Age: 88}}
	e := httpexpect.New(t, testUrl) //创建一个httpExpect实例
	e.GET("/student/redis/list").   //get请求
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().                //json body实例化
					ContainsKey("data").     //检验是否包括key
					ValueEqual("data", want) //对比key的value
}
func TestStudentController_redisDeleteOk(t *testing.T) {
	e := httpexpect.New(t, testUrl) //创建一个httpExpect实例
	e.DELETE("/student/redis/18").  //ge请求
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().              //json body实例化
					ContainsKey("ok").     //检验是否包括key
					ValueEqual("ok", true) //对比key的value
}

func TestStudentController_Postok(t *testing.T) {
	e := httpexpect.New(t, testUrl) //创建一个httpexpect实例
	//stu := &model.Student{Id: 18, Name: "555", Age: 18}
	postData := map[string]interface{}{ //创建一个json变量
		"id":   18,
		"name": "555",
		"age":  18,
	}
	contentType := "application/json;charset=utf-8"

	e.POST("/student/mysql/"). //post 请求
					WithHeader("ContentType", contentType). //定义头信息
					WithJSON(postData).                     //传入json body
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().              //json body实例化
					ContainsKey("ok").     //检验是否包括key
					ValueEqual("ok", true) //对比key的value

}

func TestStudentController_GetOk(t *testing.T) {
	stu := &model.Student{Id: 18, Name: "555", Age: 18}
	e := httpexpect.New(t, testUrl) //创建一个httpExpect实例
	e.GET("/student/mysql/18").     //get请求
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().               //json body实例化
					ContainsKey("data").    //检验是否包括key
					ValueEqual("data", stu) //对比key的value

}
func TestStudentController_GetFail(t *testing.T) {
	e := httpexpect.New(t, testUrl) //创建一个httpExpect实例
	e.GET("/student/mysql/88").     //ge请求
					Expect().
					Status(http.StatusNotFound). //判断请求是否404
					JSON().
					Object().                  //json body实例化
					ContainsKey("data").       //检验是否包括key
					ValueEqual("data", "null") //对比key的value
}

func TestStudentController_Putok(t *testing.T) {

	e := httpexpect.New(t, testUrl) //创建一个httpexpect实例
	//stu := &model.Student{Id: 18, Name: "555", Age: 18}
	postData := map[string]interface{}{ //创建一个json变量
		"name": "xiao",
		"age":  88,
	}
	contentType := "application/json;charset=utf-8"

	e.PUT("/student/mysql/18"). //put 请求
					WithHeader("ContentType", contentType). //定义头信息
					WithJSON(postData).                     //传入json body
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().              //json body实例化
					ContainsKey("ok").     //检验是否包括key
					ValueEqual("ok", true) //对比key的value

}

func TestStudentController_GetList(t *testing.T) {
	want := []*model.Student{{Id: 18, Name: "xiao", Age: 88}}
	e := httpexpect.New(t, testUrl) //创建一个httpExpect实例
	e.GET("/student/mysql/list").   //get请求
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().                //json body实例化
					ContainsKey("data").     //检验是否包括key
					ValueEqual("data", want) //对比key的value
}
func TestStudentController_deleteOk(t *testing.T) {
	e := httpexpect.New(t, testUrl) //创建一个httpExpect实例
	e.DELETE("/student/mysql/18").  //ge请求
					Expect().
					Status(http.StatusOK). //判断请求是否200
					JSON().
					Object().              //json body实例化
					ContainsKey("ok").     //检验是否包括key
					ValueEqual("ok", true) //对比key的value
}
