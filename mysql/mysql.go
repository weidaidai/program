package main

import (
	"database/sql" //database/sql仅提供基本的接口，还需指定一个第三方的数据库
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个初始化数据库的函数
func openDB(dsn string) (*sql.DB, error) {
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

type Student struct {
	Id   int
	Name string
	Age  int
}

// TODO 删除STUDENT表

func dropTable(db *sql.DB) error {
	sql := "drop table Student"
	if _, err := db.Exec(sql); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Table drop successfully")

	return nil
}

// TODO 创建STUDENT表
func createStudentTable(db *sql.DB) error {

	TABLE := `
            CREATE TABLE Student (
                Id  INT AUTO_INCREMENT,
                Name VARCHAR(50) ,
                Age INT ,
                PRIMARY KEY (Id)
            );`
	if _, err := db.Exec(TABLE); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Table created successfully")

	return nil
}

// TODO 保存Student
func saveStudent(db *sql.DB, s *Student) error {
	sql := "insert into Student(NAME,AGE)values(?,?)"

	result, err := db.Exec(sql, &s.Name, &s.Age)
	if err != nil {
		return err
	}
	newid, _ := result.LastInsertId()
	if err != nil {
		fmt.Println("seve failed", err)
		return err
	}
	fmt.Printf("new id为%d\n", newid)
	return nil

}

// TODO 删除Student
func deleteStudent(db *sql.DB, id int) error {
	sql := "delete from Student where Id=?"
	result, err := db.Exec(sql, id)
	if err != nil {
		fmt.Println("delete failed", err)
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		fmt.Println("delete  student failed", err)
		return err
	}

	fmt.Printf("delete  student successful：%d\n", id)
	return nil

}

// TODO 根据ID查Student
func getStudentById(db *sql.DB, id int) (*Student, error) {
	sql := "select ID, NAME, AGE from Student where ID=?"
	rows, e := db.Query(sql, id)
	if e == nil {
		errors.New("Query err")
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var s Student
		err := rows.Scan(&s.Id, &s.Name, &s.Age)
		if err != nil {
			return nil, err
		}
		fmt.Printf("id:%d name:%s age:%d\n", s.Id, s.Name, s.Age)
	}
	return nil, nil
}

// TODO 更新Student
func updateStudent(db *sql.DB, s *Student) error {
	sql_syntax := "update Student set Name=? where Id=?"
	result, err := db.Exec(sql_syntax, s.Name, s.Id)
	if err != nil {
		fmt.Println("Exec update student failed", err)
		return err
	}
	var id int64
	id, err = result.RowsAffected()
	if err != nil {
		fmt.Println("update student failed", err)
		return err
	}

	fmt.Printf("update student successful：%d\n", id)
	return nil
}

// TODO 把所有Student拿出来
func listAllStudents(db *sql.DB) ([]*Student, error) {
	sql := "select *from Student"

	rows, err := db.Query(sql, 0)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	//循环读取结果
	for rows.Next() {
		var s Student
		//将每一行的结果都赋值到一个s对象中
		err := rows.Scan(&s.Id, &s.Name, &s.Age)
		if err != nil {
			fmt.Println("rows failed", err)
			return nil, err
		}
		fmt.Printf("id:%d name:%s age:%d\n", s.Id, s.Name, s.Age)

	}

	return nil, nil
}

func main() {

	db, err := openDB("root:123456@tcp(127.0.0.1:3306)/sql_domo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection successful")
	//释放资源
	defer db.Close()
	//createStudentTable(db)
	//s:= &Student{ Name: "weidongqi", Age: 22}
	//saveStudent(db,s)
	//getStudentById(db, 3)
	//s:= &Student{Id: 2, Name: "weidongqi", Age: 2}
	//updateStudent(db,s)
	//listAllStudents(db)
	//deleteStudent(db,1)
	//dropTable(db)

}
