package main

import (
	"program/config"
	"program/controller"
	"program/database"

	"github.com/go-redis/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
		PoolSize: 100,
	})
	err := config.Openclient(rdb)
	if err != nil {
		return
	}
	r := gin.Default()
	db, err := config.OpenDB("root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	defer db.Close()
	defer rdb.Close()
	my_redis := database.NewredisStudentService(rdb)
	mysql_svc := database.NewMySqlStudentService(db)
	newstudent := controller.New(mysql_svc, my_redis)
	newstudent.Studentrouter(r)
	err = r.Run()
	if err != nil {
		panic(err)
	}

}
