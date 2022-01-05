package database

import (
	"errors"
	"fmt"
	"program/model"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/mitchellh/mapstructure"
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
	if svc.StuExist(key) != 1 {
		return errors.New("no exist  ")
	}
	err := svc.redis.Del(key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (svc *redisStudentService) GetStudent(id int) (*model.Student, error) {
	key := "std:" + strconv.Itoa(id)
	stu := &model.Student{}

	if svc.StuExist(key) != 1 {
		return nil, errors.New("no exist")
	}
	val, err := svc.redis.HGetAll(key).Result()
	if err != nil {
		return nil, err
	}
	//map 转结构体无法识别int
	err2 := mapstructure.Decode(val, &stu)
	if err2 != nil {
		err2.Error()
	}

	return stu, nil
}

func (svc *redisStudentService) ListStudents() ([]*model.Student, error) {
	results := make([]*model.Student, 0, 10)
	var cursor uint64

	for cursor > 0 {
		stu := &model.Student{}
		var keys []string
		keys, cursor, err := svc.redis.Scan(cursor, "stu*", 10).Result()
		if err != nil {
			return nil, err
		}
		for _, key := range keys {
			feild, err := svc.redis.HGetAll(key).Result()
			if err != nil {
				return nil, err
			}
			err2 := mapstructure.Decode(feild, &stu)
			if err2 != nil {
				err2.Error()
			}
			results = append(results, stu)
		}
		if cursor == 0 {
			break
		}
	}
	fmt.Println(results)
	return results, nil
}
