package main

import (
	"program/Config"
	"program/routers"
)

func main() {
	db, err := Config.OpenDB("root:123456@tcp(127.0.0.1:3306)/sql_domo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}

	//释放资源
	defer db.Close()
	r := routers.Router()
	r.Run()

}
