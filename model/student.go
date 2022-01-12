package model

type Student struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name"form:"name"`
	Age  int    `json:"age"form:"age"`
}
