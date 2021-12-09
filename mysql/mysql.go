package main

import (
	"database/sql" //database/sql仅提供基本的接口，还需指定一个第三方的数据库
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// User   用户对象，定义了用户的基础信息
type user struct {
	id   int    //用户id
	age  int    // 年龄
	name string //姓名
}

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_domo"
	//初始化全局的db对象
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Println("连接MYSQL失败", err)
		return err
	}

	return nil
}

//单行查询
func queryrow() {
	single_data := "select id,name,age from user where id=?"
	var u user
	err := db.QueryRow(single_data, 0).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Println("查询失败", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d", u.id, u.name, u.age)

}

// 查询多条数据
func queryMultiRow() {
	sql_syntax := "select id, name, age from user where id > ?"
	rows, err := db.Query(sql_syntax, 0)
	if err != nil {
		fmt.Println("查询多条数据失败", err)
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

//插入
func insertrow() {
	sql_syntax := "insert into user(name,age)values(?,?)"
	result, err := db.Exec(sql_syntax, "小微", 22)
	if err != nil {
		fmt.Println("插入数据失败", err)
		return
	}
	newid, err := result.LastInsertId()
	if err != nil {
		fmt.Println("获取插入的id失败", err)
		return
	}
	fmt.Printf("新插入的id为%d\n", newid)

}

//更新
func updaterow() {
	sql_syntax := "update user set name=? where id=?"
	result, err := db.Exec(sql_syntax, "小可爱", 2)
	if err != nil {
		fmt.Println("插入数据失败", err)
		return
	}
	var n int64
	n, err = result.RowsAffected()
	if err != nil {
		fmt.Println("更新数据失败", err)
		return
	}
	fmt.Printf("更新行数为：%d\n", n)

}

//删除
func deleterow() {
	sql_syntax := "delete from user where id=?"
	result, err := db.Exec(sql_syntax, 1)
	if err != nil {
		fmt.Println("删除失败", err)
		return
	}

	delid, err := result.RowsAffected()
	if err != nil {
		fmt.Println("删除失败", err)
		return
	}
	fmt.Printf("删除id在：%d\n", delid)

}
func main() {
	if err := initDB(); err != nil {
		fmt.Println("初始化连接失败", err)
		return
	}
	fmt.Println("连接成功")
	//释放资源
	defer db.Close()
	initDB()
	//queryrow()
	queryMultiRow()
	//insertrow()
	updaterow()
	//deleterow()

}
