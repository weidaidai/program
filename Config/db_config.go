package Config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// 定义一个初始化数据库的函数
func OpenDB(dsn string) (*sql.DB, error) {
	//初始化全局的db对象
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(5)
	// 尝试与数据库建立连接（校验dsn是否正确）
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
