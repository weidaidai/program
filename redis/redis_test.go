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
		DB:       1,
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
				Addr:     "127.0.0.1:63799",
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

//testing get complete
func Test_redis_get(t *testing.T) {
	//连接数据库
	rdb := preparerdb(t)
	defer rdb.Close()
	//test结束后删key
	//defer Delkey(rdb, "xiaoming")
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
			name:    "",
			args:    args{rdb: rdb, key: "xiaoming"},
			want:    &User{},
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

//testing string complete
func Test_redis_set(t *testing.T) {
	//连接数据库
	rdb := preparerdb(t)
	//延迟关闭
	defer rdb.Close()
	//test结束后删key
	//defer Delkey(rdb, "redis")
	type args struct {
		rdb *redis.Client
		u   *User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Set up the data",
			args: args{rdb: rdb, u: &User{
				Key: "redis", val: 100,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := redis_set(tt.args.rdb, tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("redis_set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


func Test_hash(t *testing.T) {
	//连接数据库
	rdb := preparerdb(t)
	//延迟关闭
	defer rdb.Close()
	//test结束后删key
	//defer Delkey(rdb,"")


	type args struct {
		rdb *redis.Client
		key string
		m   map[string]interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "hash the data",
			args: args{rdb: rdb, key: "classtwo",  m: map[string]interface{}{["id"]=1,["name"]="apple",["age"]=18,}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := hash(tt.args.rdb, tt.args.key, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("hash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
