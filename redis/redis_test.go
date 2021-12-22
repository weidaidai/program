package main

import (
	"github.com/go-redis/redis"
	"reflect"
	"testing"
)

func preparerdb(t *testing.T) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})
	err := openclient(rdb)
	if err != nil {
		t.Error(err)
	}
	return rdb
}

func Test_openclient(t *testing.T) {
	type args struct {
		rdb *redis.Client
	}
	tests := []struct {
		name    string
		args    args
		wantDb  bool
		wantErr bool
	}{
		{
			name: "good case",
			args: args{redis.NewClient(&redis.Options{
				Addr:     "127.0.0.1:6379",
				Password: "",
				DB:       0,
				PoolSize: 100,
			})},
			wantDb:  true,
			wantErr: false,
		},
		{
			name: "bad case",
			args: args{redis.NewClient(&redis.Options{
				Addr:     "127.0..q0.1:6379",
				Password: "",
				DB:       0,
				PoolSize: 100,
			})},
			wantDb:  false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := openclient(tt.args.rdb); (err != nil) != tt.wantErr {
				t.Errorf("openclient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//get
func Test_redis_get(t *testing.T) {
	//连接数据库
	rdb := preparerdb(t)
	defer rdb.Close()
	//test结束后删表
	defer Delkey(rdb, "xiaoming")
	//插入数据
	u := &User{Key: "xiaoming", val: "100", time: 0}
	redis_set(rdb, u)

	type args struct {
		rdb *redis.Client
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "",
			args: args{rdb: rdb, key: "xiaoming"},
			want: &User{
				Key: "xiaoming", val: "100",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := redis_get(tt.args.rdb, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("redis_get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("redis_get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
