package main

import (
	"database/sql" //database/sql仅提供基本的接口，还需指定一个第三方的数据库
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
	sqlStr := "insert into Student(Id,NAME,Age)values(?,?,?)"

	_, err := db.Exec(sqlStr, &s.Id, &s.Name, &s.Age)
	if err == sql.ErrNoRows {
		panic(err)
	} else {
		fmt.Println(s.Id, s.Name, s.Age)
	}

	return err
}

// TODO 删除Student
func deleteStudent(db *sql.DB, id int) error {
	sqlStr := "delete from Student where Id=?"
	_, err := db.Exec(sqlStr, id)
	if err == sql.ErrNoRows {
		panic(err)
	} else {
		fmt.Println(id)
	}
	return err

}

// TODO 更新Student
func updateStudent(db *sql.DB, s *Student) error {

	sqlStr := "update Student set Name=?,age=? where Id=?"

	row, err := db.Exec(sqlStr, &s.Name, &s.Age, &s.Id)
	if err != nil {
		return err
	}
	if r, err := row.RowsAffected(); err == nil {
		if r <= 0 {
			return err
		}
		fmt.Println(r)

	}
	return err
}

// TODO 根据ID查Student
func getStudentById(db *sql.DB, id int) (*Student, error) {
	sqlStr := "select ID, NAME, AGE from Student where ID=?"
	var U = &Student{}
	err := db.QueryRow(sqlStr, id).Scan(&U.Id, &U.Name, &U.Age)
	if err == sql.ErrNoRows {
		panic(err)
	} else {
		fmt.Println(U.Id, U.Name, U.Age)
	}
	// 循环读取结果集中的数据
	return U, nil
}

// TODO 把所有Student拿出来
func listAllStudents(db *sql.DB) ([]*Student, error) {
	u := make([]*Student, 0, 10)
	sql := "select *from Student"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//循环读取结果

	for rows.Next() {

		var s Student
		//将每一行的结果都赋值到一个s对象中
		err := rows.Scan(&s.Id, &s.Name, &s.Age)
		if err != nil {

			return nil, err
		}

		fmt.Printf("id:%d name:%s age:%d\n", s.Id, s.Name, s.Age)
		u = append(u, &s)
	}

	return u, err
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
	//s1:= &Student{ Id:2,Name: "weidongqi", Age: 22}
	//saveStudent(db,s1)
	//getStudentById(db, 2)
	s := &Student{Id: 2, Name: "weidongqi", Age: 2}
	updateStudent(db, s)
	//listAllStudents(db)
	//deleteStudent(db,1)
	//dropTable(db)

}
