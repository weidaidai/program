package database

import (
	"errors"
	"fmt"
	"program/model"
	"strconv"

	"github.com/go-redis/redis"
)

type redisStudentService struct {
	redis *redis.Client
}

func (svc *redisStudentService) StuExist(key string) int64 {

	val := svc.redis.Exists(key).Val()
	if val != 1 {
		return 0
	}
	return 1
}
func (svc *redisStudentService) SaveStudent(std *model.Student) error {

	key := "std:" + strconv.Itoa(std.Id)
	if svc.StuExist(key) != 0 {
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

func (svc *redisStudentService) UpdateStudent(std *model.Student) error {
	key := "std:" + strconv.Itoa(std.Id)
	if svc.StuExist(key) != 1 {
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

	fmt.Println(vals)
	return nil, err
}

func (svc *redisStudentService) ListStudents() ([]*model.Student, error) {

	s := &model.Student{}
	U := make([]*model.Student, 0, 10)
	var cursor uint64

	for {
		var keys []string
		//*扫描所有key，每次10条
		keys, cursor, err := svc.redis.Scan(cursor, "std*", 10).Result()
		if err != nil {
			return nil, err
		}
		for _, key := range keys {
			_, err := svc.redis.HGetAll(key).Result()
			if err != nil {
				return nil, err
			}
			//在range
			U = append(U, s)
		}

		if cursor == 0 {
			break
		}
	}
	return U, nil

}
