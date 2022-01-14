package database

import (
	"database/sql"
	"errors"
	"fmt"
	"program/model"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlStudentService struct {
	db *sql.DB
}

func NewMySqlStudentService(db *sql.DB) StudentService {
	return &mysqlStudentService{db: db}
}

func (svc *mysqlStudentService) SaveStudent(std *model.Student) error {

	sqlStr := "INSERT INTO STUDENT(ID,NAME,AGE)VALUES (?,?,?)"
	_, err := svc.db.Exec(sqlStr, &std.Id, &std.Name, &std.Age)
	return err

}

func (svc *mysqlStudentService) UpdateStudent(id int, std *model.Student) error {
	sqlStr := "UPDATE STUDENT SET NAME=?,AGE=? where ID=?"
	row, err := svc.db.Exec(sqlStr, &std.Name, &std.Age, id)
	if err != nil {
		return err
	}
	r, _ := row.RowsAffected()
	if r <= 0 {
		return errors.New("Update the same data or data does not exist")
	}
	return err
}

func (svc *mysqlStudentService) DeleteStudent(id int) error {
	sqlStr := "DELETE FROM STUDENT WHERE ID=?"
	_, err := svc.db.Exec(sqlStr, id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *mysqlStudentService) GetStudent(id int) (*model.Student, error) {

	sqlStr := "SELECT ID, NAME, AGE FROM STUDENT WHERE ID=?"
	stu := &model.Student{}
	err := svc.db.QueryRow(sqlStr, id).Scan(&stu.Id, &stu.Name, &stu.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return stu, nil
}

func (svc *mysqlStudentService) ListStudents() ([]*model.Student, error) {
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
func (svc *mysqlStudentService) createStudentTable() error {
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
func (svc *mysqlStudentService) dropTable() error {
	sql := "DROP TABLE STUDENT "
	if _, err := svc.db.Exec(sql); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
