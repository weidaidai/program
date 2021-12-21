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

func Test_redis_set(t *testing.T) {
	//连接数据库
	rdb := preparerdb(t)
	defer rdb.Close()

	u := &User{Key: "test1", val: "val", time: 8}
	redis_set(rdb, u)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := redis_set(tt.args.rdb, tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("redis_set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redis_get(t *testing.T) {
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
		// TODO: Add test cases.
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
