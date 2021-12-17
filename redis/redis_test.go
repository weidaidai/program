package main

import (
	"github.com/go-redis/redis"
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
			name:    "good case",
			args:    args{redis.NewClient(&redis.Options{
				Addr:     "127.0.0.1:6379",
				Password: "",
				DB:       0,
				PoolSize: 100,
			})},
			wantDb:  true,
			wantErr: false
		},
		{
			name:    "bad case",
			args:    args{redis.NewClient(&redis.Options{
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

func Test_hash(t *testing.T) {
	rdb:=preparerdb(t)
	defer rdb.Close()
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}