package main

import (
	"os"
	"program/config"
	"program/controller"
	"program/database"
	"strconv"

	"github.com/go-redis/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	//dns=root:123456@tcp(127.0.0.1:33306)/sql_test?charset=utf8&parseTime=True&loc=Local

	MysqlDNS := os.Getenv("MYSQL_DNS")
	RedisAddr := os.Getenv("REDIS_ADDR")
	Redis_Passworld := os.Getenv("REDIS_PASSWORD")
	RedisPoolsize, err := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		PoolSize: RedisPoolsize,
		Password: Redis_Passworld,
	})

	defer rdb.Close()
	db, err := config.OpenDB(MysqlDNS)
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
