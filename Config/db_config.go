package Config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB(dns string) (*sql.DB, error) {
	//初始化全局的db对象

	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(10)

	db.SetMaxIdleConns(5)
	//ping
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
