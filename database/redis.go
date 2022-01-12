package database

import (
	"errors"
	"program/model"
	"strconv"

	"github.com/go-redis/redis"
)

type RedisStudentService struct {
	redis *redis.Client
}

func (svc *RedisStudentService) stuExist(key string) bool {
	val := svc.redis.Exists(key).Val()
	return val == 1

}

func (svc *RedisStudentService) SaveStudent(std *model.Student) error {

	key := "std:" + strconv.Itoa(std.Id)
	if svc.stuExist(key) {
		return errors.New("id repetition ")
	}

	statusCmd := svc.redis.HMSet(key, map[string]interface{}{
		"id":   std.Id,
		"age":  std.Age,
		"name": std.Name,
	})

	if err := statusCmd.Err(); err != nil {
		return err
	}

	return nil
}

func (svc *RedisStudentService) UpdateStudent(std *model.Student) error {
	key := "std:" + strconv.Itoa(std.Id)
	if !svc.stuExist(key) {
		return errors.New("no exist  ")
	}

	statusCmd := svc.redis.HMSet(key, map[string]interface{}{
		"id":   std.Id,
		"age":  std.Age,
		"name": std.Name,
	})

	if err := statusCmd.Err(); err != nil {
		return err
	}
	return nil
}

func (svc *RedisStudentService) DeleteStudent(id int) error {
	key := "std:" + strconv.Itoa(id)
	err := svc.redis.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}
func (svc *RedisStudentService) getStudentbykey(key string) (*model.Student, error) {
	var id int
	key = "std:" + strconv.Itoa(id)
	return svc.GetStudent(id)
}
func (svc *RedisStudentService) GetStudent(id int) (*model.Student, error) {
	key := "std:" + strconv.Itoa(id)
	val, err := svc.redis.HGetAll(key).Result()
	if err != nil {
		return nil, nil
	}
	stu := &model.Student{}
	stu.Id, err = strconv.Atoi(val["id"])
	if err != nil {
		return nil, err
	}
	stu.Name = val["name"]
	stu.Age, err = strconv.Atoi(val["age"])
	if err != nil {
		return nil, err
	}

	return stu, nil
}

func (svc *RedisStudentService) ListStudents() ([]*model.Student, error) {

	results := make([]*model.Student, 0, 10)
	var cursor uint64

	for {
		keys, cursor, err := svc.redis.Scan(cursor, "std:*", 10).Result()
		if err != nil {
			return nil, err
		}
		for _, key := range keys {
			std, err := svc.getStudentbykey(key)
			if err != nil {
				return nil, err
			}
			results = append(results, std)

		}
		if cursor == 0 {
			break
		}

	}

	return results, nil

}
