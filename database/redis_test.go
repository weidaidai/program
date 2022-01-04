package database

import (
	"program/model"
	"reflect"
	"strconv"
	"testing"

	"github.com/go-redis/redis"
)

func preparerdb(t *testing.T) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
		PoolSize: 100,
	})

	return rdb
}
func insert(t *testing.T, rdb *redis.Client, std *model.Student) {
	key := "std:" + strconv.Itoa(std.Id)
	statusCmd := rdb.HMSet(key, map[string]interface{}{
		"id":   std.Id,
		"age":  std.Age,
		"name": std.Name,
	})
	if err := statusCmd.Err(); err != nil {
		t.Error()
	}
}
func Test_redisStudentService_SaveStudent(t *testing.T) {

	t.Run("save not exist", func(t *testing.T) {
		rdb := preparerdb(t)
		defer rdb.Close()
		defer rdb.FlushAll()
		s := &model.Student{
			Id:   1,
			Name: "xiao",
			Age:  18,
		}
		wantErr := false
		svc := &redisStudentService{
			redis: rdb,
		}
		if err := svc.SaveStudent(s); (err != nil) != wantErr {
			t.Errorf("saveStudent() error = %v, wantErr %v", err, wantErr)

		}

	})

	t.Run("save exist", func(t *testing.T) {
		rdb := preparerdb(t)
		defer rdb.FlushAll()
		s1 := &model.Student{Id: 5, Name: "weidongqi", Age: 22}
		insert(t, rdb, s1)
		s := &model.Student{
			Id:   5,
			Name: "xiao",
			Age:  18,
		}

		wantErr := true
		svc := &redisStudentService{
			redis: rdb,
		}
		if err := svc.SaveStudent(s); (err != nil) != wantErr {
			t.Errorf("saveStudent() error = %v, wantErr %v", err, wantErr)
		}

	})

}

func Test_redisStudentService_UpdateStudent(t *testing.T) {

	t.Run("update not exist", func(t *testing.T) {
		rdb := preparerdb(t)
		defer rdb.Close()
		defer rdb.Del()
		s := &model.Student{
			Id:   9,
			Name: "xiao",
			Age:  18,
		}
		wantErr := true
		svc := &redisStudentService{
			redis: rdb,
		}
		if err := svc.UpdateStudent(s); (err != nil) != wantErr {
			t.Errorf("saveStudent() error = %v, wantErr %v", err, wantErr)
		}

	})

	t.Run("updata exist", func(t *testing.T) {
		rdb := preparerdb(t)
		defer rdb.Close()
		defer rdb.FlushAll()
		s1 := &model.Student{Id: 5, Name: "weidongqi", Age: 22}
		insert(t, rdb, s1)
		s := &model.Student{Id: 5, Name: "xiaoming", Age: 18}

		wantErr := false
		svc := &redisStudentService{
			redis: rdb,
		}
		if err := svc.UpdateStudent(s); (err != nil) != wantErr {
			t.Errorf("saveStudent() error = %v, wantErr %v", err, wantErr)
		}

	})
}

func Test_redisStudentService_DeleteStudent(t *testing.T) {
	type fields struct {
		redis *redis.Client
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "del not exist",
			fields: fields{
				redis: preparerdb(t)},
			args: args{
				id: 5,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &redisStudentService{
				redis: tt.fields.redis,
			}
			if err := svc.DeleteStudent(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisStudentService_GetStudent(t *testing.T) {

	t.Run("get exist", func(t *testing.T) {
		rdb := preparerdb(t)
		defer rdb.Close()
		//defer rdb.FlushAll()
		s := &model.Student{Id: 1, Name: "xiaoxiaoxing", Age: 22}
		insert(t, rdb, s)

		wantErr := false
		svc := &redisStudentService{
			redis: rdb,
		}
		s2, err2 := svc.GetStudent(1)
		if err2 != nil {
			t.Errorf("saveStudent() error = %v, wantErr %v", err2, wantErr)
		}
		if !reflect.DeepEqual(s2, s) {
			t.Errorf("getStudentById() got = %v, want %v", s2, s)
		}
	})
}

func Test_redisStudentService_ListStudents(t *testing.T) {

}
