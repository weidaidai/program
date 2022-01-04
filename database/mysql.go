package database

import (
	"database/sql"
	"errors"
	"fmt"
	"program/model"

	_ "github.com/go-sql-driver/mysql"
)

func dropTable(db *sql.DB) error {
	sql := "drop table Student"
	if _, err := db.Exec(sql); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

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

	return nil
}

type MysqlStudentService struct {
	db *sql.DB
}

func (svc *MysqlStudentService) SaveStudent(std *model.Student) error {

	sqlStr := "insert into Student(Id,NAME,Age)values(?,?,?)"
	_, err := svc.db.Exec(sqlStr, &std.Id, &std.Name, &std.Age)
	if err == sql.ErrNoRows {
		errors.New("save failed")
	}
	return err

}

func (svc *MysqlStudentService) UpdateStudent(std *model.Student) error {
	sqlStr := "update Student set Name=?,age=? where Id=?"
	row, err := svc.db.Exec(sqlStr, &std.Name, &std.Age, &std.Id)
	if err != nil {
		return err
	}
	r, _ := row.RowsAffected()
	if r <= 0 {
		errors.New("Update the same data or data does not exist")
	}
	return err
}

func (svc *MysqlStudentService) DeleteStudent(id int) error {
	sqlStr := "delete from Student where Id=?"
	row, err := svc.db.Exec(sqlStr, id)

	r, _ := row.RowsAffected()
	if r <= 0 {
		//panic(err)
		errors.New("NO data delete")
	}
	return err
}

func (svc *MysqlStudentService) GetStudent(id int) (*model.Student, error) {
	sqlStr := "select ID, NAME, AGE from Student where ID=?"
	U := &model.Student{}
	//var U *model.Student 尽量简短声明
	err := svc.db.QueryRow(sqlStr, id).Scan(&U.Id, &U.Name, &U.Age)
	if err == sql.ErrNoRows {
		errors.New("no row get")
	}
	// 循环读取结果集中的数据
	return U, nil
}

func (svc *MysqlStudentService) ListStudents() ([]*model.Student, error) {
	u := make([]*model.Student, 0, 10)
	sql := "select *from Student"
	rows, err := svc.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//循环读取结果

	for rows.Next() {

		s := &model.Student{}

		err := rows.Scan(&s.Id, &s.Name, &s.Age)
		if err != nil {

			return nil, err
		}
		fmt.Printf("id:%d name:%s age:%d\n", s.Id, s.Name, s.Age)
		//将每一行的结果都赋值到一个u对象中
		u = append(u, s)
	}
	return u, err
}
