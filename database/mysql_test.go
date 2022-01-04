package database

import (
	"database/sql"
	"program/Config"
	"testing"
)

func prepareDB(t *testing.T) *sql.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err := Config.OpenDB(dsn)
	if err != nil {
		t.Error(err)
	}
	return db
}

func prepareTable(t *testing.T, db *sql.DB) {
	err := createStudentTable(db)
	if err != nil {
		t.Error(err)
	}
}

func TestMysqlStudentService_DeleteStudent(t *testing.T) {

}

func TestMysqlStudentService_GetStudent(t *testing.T) {

}

func TestMysqlStudentService_ListStudents(t *testing.T) {

}

func TestMysqlStudentService_SaveStudent(t *testing.T) {

}

func TestMysqlStudentService_UpdateStudent(t *testing.T) {

}
