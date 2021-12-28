package service

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"program/model"
	"strconv"
)

type redisStudentService struct {
	redis *redis.Client
}

func (svc *redisStudentService) SaveStudent(std *model.Student) error {
	// TODO ID不能重复，重复要error
	key := "std:" + strconv.Itoa(std.Id)
	if strconv.Itoa(std.Id) == svc.redis.HGet(key, "id").Val() {
		return errors.New("id repetition ")
	} else {
		statusCmd := svc.redis.HMSet(key, map[string]interface{}{
			"id":   std.Id,
			"age":  std.Age,
			"name": std.Name,
		})

		if err := statusCmd.Err(); err != nil {
			return err
		}
	}

	return nil
}

func (svc *redisStudentService) UpdateStudent(std *model.Student) error {
	// TODO 只更新存在的，不存在要error
	key := "std:" + strconv.Itoa(std.Id)
	err := svc.redis.HGet(key, "id").Err()
	if err != nil {
		return err
	} else {
		statusCmd := svc.redis.HMSet(key, map[string]interface{}{
			"id":   std.Id,
			"age":  std.Age,
			"name": std.Name,
		})

		if err := statusCmd.Err(); err != nil {
			return err
		}
	}
	return nil
}

func (svc *redisStudentService) DeleteStudent(id int) error {
	key := "std:" + strconv.Itoa(id)
	num := svc.redis.Del(key).Val()
	if num == 0 {
		errors.New("del failed")
	}
	return nil
}

func (svc *redisStudentService) GetStudent(id int) (map[string]interface{}, error) {
	key := "std:" + strconv.Itoa(id)
	vals, err := svc.redis.HMGet(key, "id", "name", "age").Result()
	if err != nil {
		return nil, err
	}
	// vals是一个数组
	fmt.Println(vals)
	return nil, err
}

func (svc *redisStudentService) ListStudents() ([]*model.Student, error) {
	// TODO 用 SCAN
	num := make([]*model.Student, 0, 10)
	val := svc.redis.Scan(0, "*std", 10).Args()
	//svc.redis.Keys("*").Err()
	stu := model.Student{}
	num = append(num, &stu)
	fmt.Println(val)
	return num, nil

}

func openclient(rdb *redis.Client) (err error) {

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return err
}
