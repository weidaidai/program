package config

import (
	"database/sql"

	"github.com/go-redis/redis"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(10)
	db.SetMaxIdleConns(5)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
func Openclient(rdb *redis.Client) (err error) {

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return err
}
