package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis" //自带原生连接池
)


// init

func initclient() (rdb *redis.Client,err error) {

	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return rdb,err
}

//string
func redis_set(rdb *redis.Client) {
	err := rdb.Set("name", 1000, 0).Err()
	if err != nil {
		log.Fatal("无法设置", err)
		return
	}
	var val string
	val, err = rdb.Get("name").Result()
	if err != nil {
		fmt.Println("获取值失败", err)
		return
	}
	fmt.Println("name", val)
}
func setex(rdb *redis.Client) {
	err := rdb.Set("set1", 188, 0).Err()
	if err != nil {
		fmt.Println("set设置错误", err)
		return
	}

	val2, err := rdb.SMembers("set").Result()
	if err != nil {
		fmt.Println("获取set值失败", err)
		return

	}
	fmt.Println("set1", val2)
}

//hash
func hash(rdb *redis.Client) {

	data := make(map[string]interface{})
	data["id"] = 1
	data["name"] = "小明"
	data["age"] = 18
	// 一次性保存多个hash字段值
	err := rdb.HMSet("class2", data).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("hash 设置成功")

	// HMGet支持多个field字段名，意思是一次返回多个字段值
	vals, err := rdb.HMGet("class2", "id", "name").Result()
	if err != nil {
		panic(err)
	}

	// vals是一个数组
	fmt.Println(vals)

}

//list
func list(rdb *redis.Client) {
	err := rdb.LPush("list1", "redis").Err()
	if err != nil {
		panic(err)

	}
	fmt.Println("list 设置成功")
	vals, err := rdb.LRange("list1", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)

}

//set

func set(rdb *redis.Client) {
	err := rdb.SAdd("set1", "mysql").Err()
	if err != nil {
		panic(err)

	}
	fmt.Println("set 设置成功")
	vals, err := rdb.SMembers("set1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)

}

//zset
func zset(rdb *redis.Client) {
	zsetkey := "database"
	languages := []redis.Z{
		redis.Z{95, "mysql"},
		redis.Z{66, "redis1"},
		redis.Z{77, "mogondb1"},
	}
	num, err := rdb.ZAdd(zsetkey, languages...).Result()
	if err != nil {
		log.Fatal("zdd 错误",err)
		return
	}
	fmt.Println("值为", num)
}

//zrange
func zranges(rdb *redis.Client) {

	vals, err := rdb.ZRange("database", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("database 结果为", vals)
}

func main() {

	rdb,err:= initclient()
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功")
	//记得释放资源
	defer rdb.Close()
	redis_set(rdb)
	//setex()
	//zset()
	//hash()
	//list()
	zranges(rdb)
}
