package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

// init

func initclient() (err error) {

	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
//string
func redisset() {
	err := rdb.Set("name", 1000, 0).Err()
	if err != nil {
		fmt.Println("无法设置", err)
		return
	}

	val,err:= rdb.Get("name").Result()
	if err != nil {
		fmt.Println("获取值失败", err)
		return
	}
	fmt.Println("name", val)
}
func setex() {
	err :=rdb.Set("set1",188,0).Err()
	if err != nil {
		fmt.Println("set设置错误",err)
		return
	}

	val2,err:=rdb.SMembers("set").Result()
	if err!= nil {
		fmt.Println("获取set值失败",err)
		return

	}
	fmt.Println("set1",val2)
}
func zsent() {
	zsetkey:="yulian"
	languages:=[]redis.Z{
		redis.Z{95,"mysql"},
		redis.Z{66,"redis1"},
		redis.Z{77,"mogondb1"},

	}
	num,err:=rdb.ZAdd(zsetkey,languages...).Result()
	if err != nil {
		fmt.Println("zdd 错误")
		return
	}
	fmt.Println("值为",num)
}

func main() {
	if err := initclient();
		err != nil {

		fmt.Println("init redis faild,err:%v", err)
		return

	}
	fmt.Println("成功")

	defer rdb.Close()
	redisset()
	setex()
	zsent()
}
