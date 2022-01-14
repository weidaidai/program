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

	defer rdb.Close()
	db, err := config.OpenDB("root:123456@tcp(127.0.0.1:3306)/sql_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	redis := database.NewRedisStudentService(rdb)
	svcRedis := controller.New(redis)
	mysql := database.NewMySqlStudentService(db)
	svcMysql := controller.New(mysql)

	r := gin.Default()
	svcMysql.StudentRouter(r, "/mysql")
	svcRedis.StudentRouter(r, "/redis")
	err = r.Run()
	if err != nil {
		panic(err)
	}

}
