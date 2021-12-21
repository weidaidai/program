package main

import (
	"fmt"
	"github.com/go-redis/redis" //自带原生连接池
	"log"
	"time"
)

//  connention redis

func openclient(rdb *redis.Client) (err error) {

	_, err = rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return err
}

type User struct {
	Key  string
	val  interface{}
	time time.Duration
}

//string
func redis_set(rdb *redis.Client, u *User) error {
	err := rdb.Set(u.Key, u.val, u.time).Err()
	if err != nil {
		return err
	}
	fmt.Println("string set up successfully")
	return err
}

//complete get
func redis_get(rdb *redis.Client, key string) (*User, error) {

	val, err := rdb.Get(key).Result()
	if err != nil {
		return nil, err
	}
	var u = &User{}
	fmt.Println(key, val)
	return u, err
}

//complete hash
func hash(rdb *redis.Client, key string, m map[string]interface{}) error {
	// 一次性保存多个hash字段值
	err := rdb.HMSet(key, m).Err()
	if err != nil {
		return err
	}
	fmt.Println("Hash set up successfully")

	return err

}

//complete hash HMGet
func gethash(rdb *redis.Client, key string) (map[string]interface{}, error) {
	// HMGet支持多个field字段名，意思是一次返回多个字段值
	vals, err := rdb.HMGet(key, "name", "id", "age").Result()
	if err != nil {
		return nil, err
	}

	// vals是一个数组
	fmt.Println(vals)
	return nil, err
}

// complete list

type liststudent struct {
	key string
	val interface{}
}

func list(rdb *redis.Client, l *liststudent) error {
	num, err := rdb.LPush(l.key, l.val).Result()
	if err != nil {
		return err
	}
	if num <= 0 {
		log.Fatalln("list failed")
	} else if num > 0 {
		fmt.Println("list set up successfully ")
	}

	return err
}
func lrang(rdb *redis.Client, Key string, start int64, stop int64) ([]*User, error) {
	u := make([]*User, 0, 100)
	vals, err := rdb.LRange(Key, start, stop).Result()
	if err != nil {
		return nil, err
	}

	var s User
	u = append(u, &s)
	fmt.Println(Key, vals)
	return u, err
}

// TODO set

type set_student struct {
	key string
	val interface{}
}

func Set_student(rdb *redis.Client, student *set_student) error {
	num, err := rdb.SAdd(student.key, student.val).Result()
	if err != nil {
		return err
	}
	if num <= 0 {
		log.Fatalln("set failed")
	} else if num > 0 {
		fmt.Println("Set set up successfully")
	}

	return nil
}
func Smembers(rdb *redis.Client, key string) ([]*set_student, error) {
	s := make([]*set_student, 0, 100)
	vals, err := rdb.SMembers(key).Result()
	length, err := rdb.SCard(key).Result() //返回名称为 key 的zset的长度
	if err != nil {
		return nil, err
	}

	fmt.Println("length:", length)
	if err != nil {
		return nil, err
	}
	var s1 set_student
	s = append(s, &s1)
	fmt.Println(key, vals)
	return s, err
}

type zset_sm struct {
	Score  float64
	Member interface{}
}

//TODO zset
func zset(rdb *redis.Client, key string, l *zset_sm) error {

	num, err := rdb.ZAdd(key, redis.Z{l.Score, l.Member}).Result()
	if err != nil {
		return err
	}
	if num <= 0 {
		log.Fatalln("zset failed")
	} else if num > 0 {
		fmt.Println("zset set up successfully")
	}

	return err
}

//zrange
func zranges(rdb *redis.Client, key string, start int64, stop int64) ([]*zset_sm, error) {
	u := make([]*zset_sm, 0, 20)
	vals, err := rdb.ZRangeWithScores(key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	var s zset_sm
	u = append(u, &s)
	fmt.Println(key, vals)
	return u, err
}

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
		PoolSize: 100,
	})
	err := openclient(rdb)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection successful")
	//记得释放资源
	defer rdb.Close()
	//redis_get(rdb,"kkk")
	//u:=&User{Key: "kkk",val: 555,time: 0}
	//redis_set(rdb,u)
	//gethash(rdb,"class")
	//zset()
	//var hashdata = make(map[string]interface{})
	//hashdata  = make(map[string]interface{})
	//hashdata ["id"] = 1
	//hashdata ["name"] = "小明"
	//hashdata ["age"] = 18
	//hash(rdb,"class",hashdata)
	//L:=&liststudent{key:"lst",val: "redis"}
	//list(rdb,L)
	//lrang(rdb,"lst",0,3)
	//student:=&set_student{key:"settttt",val: "shan"}
	//Set_student(rdb,student)
	//Smembers(rdb,"settttt")
	zranges(rdb, "database", 1, -1)
	//zset_1:=&zset_sm{100,"数学"}
	//zset_2:=&zset_sm{90,"语文"}
	//zset(rdb,"database",zset_1)
	//zset(rdb,"database",zset_2)

}
