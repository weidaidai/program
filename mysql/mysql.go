package main

import (
	"database/sql" //database/sql仅提供基本的接口，还需指定一个第三方的数据库
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// User   用户对象，定义了用户的基础信息
type user struct {
	id   int    //用户id
	age  int    // 年龄
	name string //姓名
}

// 定义一个初始化数据库的函数
func initDB(dsn string) (db *sql.DB, err error) {

	//初始化全局的db对象
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return

	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(5)
	// 尝试与数据库建立连接（校验dsn是否正确）
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db, err
}

// 查询数据
func queryMultiRow(db *sql.DB) {
	sql_syntax := "select id, name, age from user where id> ? and id< ?"
	rows, err := db.Query(sql_syntax, 1, 3)
	if err != nil {
		log.Fatal("查询多条数据失败", err)
		return
	}
	// 关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			log.Fatal("查询多条数据失败", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

//插入
func insertrow(db *sql.DB) {
	sql_syntax := "insert into user(name,age)values(?,?)"
	result, err := db.Exec(sql_syntax, "小微", 22)
	if err != nil {
		log.Fatal("插入数据失败", err)
		return
	}
	newid, err := result.LastInsertId()
	if err != nil {
		log.Fatal("获取插入的id失败", err)
		return
	}
	fmt.Printf("新插入的id为%d\n", newid)

}

//更新
func Updaterow(db *sql.DB,id int64)(int64,error) {
	sql_syntax := "update user set name=? where id=?"
	result, err := db.Exec(sql_syntax, "小", 2)
	if err != nil {
		log.Fatal("插入数据失败", err)
		return -1,err
	}

	id, err = result.RowsAffected()
	if err != nil {
		log.Fatal("更新数据失败", err)
		return -1 ,err
	}
	fmt.Printf("更新行数为：%d\n", id)
    return id,err
}

//删除
func deleterow(db *sql.DB) {
	sql_syntax := "delete from user where id=?"
	result, err := db.Exec(sql_syntax, 1)
	if err != nil {
		log.Fatal("删除失败", err)
		return
	}

	delid, err := result.RowsAffected()
	if err != nil {
		log.Fatal("删除失败", err)
		return
	}
	fmt.Printf("删除id在：%d\n", delid)

}
func main() {
	db, err := initDB("root:123456@tcp(127.0.0.1:3306)/sql_domo")
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功")
	//释放资源
	defer db.Close()

	//queryMultiRow(db)
	//insertrow(db)
	Updaterow(db,4)
	//deleterow(db)


}
