package main

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)
// User   用户对象，定义了用户的基础信息
type user  struct {
	id   int   //用户id
	age float32  // 年龄
	name string  //姓名
}

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_domo"//前提要创建一个数据库sql_domo
	// 不会校验账号密码是否正确
	db, err = sql.Open("mysql", dsn)//数据库的名称mysql sql.Open是database/sql本地
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect to db failed,err:%v\n",err)
		return err
	}
	//根据业务范围确定
	db.SetMaxIdleConns(10)
	return nil
}
// 查询多条数据
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%f\n", u.id, u.name, u.age)
	}
}
func main() {
	if err := initDB();
		err != nil {
		fmt.Println("初始化连接失败",err)
		return
	}
	fmt.Println("连接成功")
	defer db.Close()
	initDB()
	queryMultiRowDemo()
}

