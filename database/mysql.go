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

	sqlStr := "INSERT INTO STUDENT(ID,NAME,AGE)VALUES (?,?,?)"
	_, err := svc.db.Exec(sqlStr, &std.Id, &std.Name, &std.Age)
	return err

}

func (svc *MysqlStudentService) UpdateStudent(std *model.Student) error {
	sqlStr := "UPDATE STUDENT SET NAME=?,AGE=? where ID=?"
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
	sqlStr := "DELETE FROM STUDENT WHERE ID=?"
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
	sqlStr := "SELECT ID, NAME, AGE FROM STUDENT WHERE ID=?"
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
	sql := "SELECT * FROM STUDENT"
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

	return stu, err
}
func (svc *MysqlStudentService) createStudentTable() error {
	TABLE := `CREATE TABLE STUDENT  (
                ID  INT AUTO_INCREMENT,
                NAME VARCHAR(50) ,
                AGE INT ,
                PRIMARY KEY (ID)
            );`
	if _, err := svc.db.Exec(TABLE); err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
func (svc *MysqlStudentService) dropTable() error {
	sql := "DROP TABLE STUDENT "
	if _, err := svc.db.Exec(sql); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
