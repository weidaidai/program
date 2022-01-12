package main

import (
	"program/config"
	"program/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := config.OpenDB("root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	defer db.Close()

	new := controller.Newmysql(db)
	new.Studentrouter(r)

	err = r.Run()
	if err != nil {
		panic(err)
	}

}
