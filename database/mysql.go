package database

import (
	"database/sql"
	"errors"
	"fmt"
	"program/model"

	_ "github.com/go-sql-driver/mysql"
)

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
		return errors.New("Update the same data or data does not exist")
	}
	return err
}

func (svc *MysqlStudentService) DeleteStudent(id int) error {
	sqlStr := "delete from Student where Id=?"
	row, err := svc.db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	r, _ := row.RowsAffected()
	if r <= 0 {
		return errors.New("NO data delete")
	}
	return nil
}

func (svc *MysqlStudentService) GetStudent(id int) (*model.Student, error) {
	sqlStr := "select ID, NAME, AGE from Student where ID=?"
	stu := &model.Student{}
	//var U *model.Student 尽量简短声明
	row := svc.db.QueryRow(sqlStr, id).Scan(&stu.Id, &stu.Name, &stu.Age)
	if row == sql.ErrNoRows {
		return nil, errors.New("no row get")
	}
	fmt.Println(stu)
	return stu, nil
}

func (svc *MysqlStudentService) ListStudents() ([]*model.Student, error) {
	stu := make([]*model.Student, 0, 10)
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
		stu = append(stu, s)
	}
	//fmt.Printf("%#v", u)
	return stu, err
}
func (svc *MysqlStudentService) createStudentTable() error {
	TABLE := `
            CREATE TABLE Student (
                Id  INT AUTO_INCREMENT,
                Name VARCHAR(50) ,
                Age INT ,
                PRIMARY KEY (Id)
            );`
	if _, err := svc.db.Exec(TABLE); err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
func (svc *MysqlStudentService) dropTable() error {
	sql := "drop table Student"
	if _, err := svc.db.Exec(sql); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
