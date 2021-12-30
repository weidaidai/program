package database

import (
	"program/model"
	"reflect"
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
	err := openclient(rdb)
	if err != nil {
		t.Error(err)
	}
	return rdb
}

func Test_redisStudentService_SaveStudent(t *testing.T) {

	rdb := preparerdb(t)

	defer rdb.Close()
	defer rdb.HDel("std11")

	type fields struct {
		redis *redis.Client
	}
	type args struct {
		std *model.Student
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "save id exist",
			fields: fields{
				redis: preparerdb(t)},
			args: args{
				std: &model.Student{
					Id:   1,
					Name: "小新",
					Age:  18,
				},
			},
			wantErr: true,
		},
		{name: "save not exist",
			fields: fields{
				redis: preparerdb(t)},
			args: args{
				std: &model.Student{
					Id:   15,
					Name: "小新xing",
					Age:  18,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &redisStudentService{
				redis: tt.fields.redis,
			}
			if err := svc.SaveStudent(tt.args.std); (err != nil) != tt.wantErr {
				t.Errorf("SaveStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func Test_redisStudentService_UpdateStudent(t *testing.T) {

	type fields struct {
		redis *redis.Client
	}
	type args struct {
		std *model.Student
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Update id exist",
			fields: fields{
				redis: preparerdb(t)},
			args: args{
				std: &model.Student{
					Id:   1,
					Name: "小新xing",
					Age:  18,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &redisStudentService{
				redis: tt.fields.redis,
			}
			if err := svc.UpdateStudent(tt.args.std); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStudent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
		{name: "del",
			fields: fields{
				redis: preparerdb(t)},
			args: args{
				id: 3,
			},
			wantErr: false,
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
		want    map[string]interface{}
		wantErr bool
	}{
		{name: "get",
			fields: fields{
				redis: preparerdb(t)},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &redisStudentService{
				redis: tt.fields.redis,
			}
			got, err := svc.GetStudent(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisStudentService_ListStudents(t *testing.T) {

	type fields struct {
		redis *redis.Client
	}

	tests := []struct {
		name    string
		fields  fields
		want    []*model.Student
		wantErr bool
	}{
		{
			name: "get key 10",
			fields: fields{
				redis: preparerdb(t),
			},
			want: []*model.Student{},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &redisStudentService{
				redis: tt.fields.redis,
			}
			got, err := svc.ListStudents()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListStudents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListStudents() got = %v, want %v", got, tt.want)
			}
		})
	}
}
