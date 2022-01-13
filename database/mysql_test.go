package database

import (
	"database/sql"
	"program/config"
	"program/model"
	"reflect"
	"testing"
)

func prepareDB(t *testing.T) *sql.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_demo2"
	db, err := config.OpenDB(dsn)
	if err != nil {
		t.Error(err)
	}
	return db
}

func TestMysqlStudentService_DeleteStudent(t *testing.T) {

	t.Run("del exist", func(t *testing.T) {
		db := prepareDB(t)
		svc := &mysqlStudentService{
			db,
		}
		svc.createStudentTable()
		defer svc.dropTable()
		s := &model.Student{Id: 1, Name: "xiaoxing", Age: 18}
		svc.SaveStudent(s)
		//test结束后删表

		wantErr := false
		err := svc.DeleteStudent(1)
		if err != nil {
			t.Errorf("DeleteStudent() error = %v, wantErr %v", err, wantErr)
		}

	})
}

func TestMysqlStudentService_GetStudent(t *testing.T) {

	t.Run("get exist", func(t *testing.T) {
		db := prepareDB(t)
		svc := &mysqlStudentService{
			db,
		}
		svc.createStudentTable()
		defer svc.dropTable()
		s := &model.Student{Id: 1, Name: "xiaoxing", Age: 18}
		svc.SaveStudent(s)

		wantErr := false
		s2, err2 := svc.GetStudent(1)
		if err2 != nil {
			t.Errorf("saveStudent() error = %v, wantErr %v", err2, wantErr)
		}
		if !reflect.DeepEqual(s2, s) {
			t.Errorf("getStudentById() got = %v, want %v", s2, s)
		}
	})

}

func TestMysqlStudentService_ListStudents(t *testing.T) {

	db := prepareDB(t)
	defer db.Close()
	svc := &mysqlStudentService{
		db,
	}
	svc.createStudentTable()
	defer svc.dropTable()

	//插入数据
	s1 := &model.Student{Id: 1, Name: "xiaoxing", Age: 22}
	s2 := &model.Student{Id: 2, Name: "xiaoxing", Age: 33}
	svc.SaveStudent(s1)
	svc.SaveStudent(s2)
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    []*model.Student
		wantErr bool
	}{
		{
			name: "",
			args: args{db: db},
			want: []*model.Student{
				{Id: 1, Name: "xiaoxing", Age: 22},
				{Id: 2, Name: "xiaoxing", Age: 33},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := svc.ListStudents()
			if (err != nil) != tt.wantErr {
				t.Errorf("listAllStudents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listAllStudents() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestMysqlStudentService_SaveStudent(t *testing.T) {

	t.Run("save not exist", func(t *testing.T) {

		db := prepareDB(t)

		defer db.Close()
		svc := &mysqlStudentService{
			db,
		}
		svc.createStudentTable()
		//test结束后删表
		defer svc.dropTable()
		s := &model.Student{
			Id:   1,
			Name: "weidongqi",
			Age:  22,
		}

		wantErr := false
		if err := svc.SaveStudent(s); (err != nil) != wantErr {
			t.Errorf("saveStudent() error = %v, wantErr %v", err, wantErr)
		}
		s2, err2 := svc.GetStudent(1)
		if err2 != nil {
			t.Error(err2)
		}
		if !reflect.DeepEqual(s2, s) {
			t.Errorf("getStudentById() got = %v, want %v", s2, s)
		}
	})

	t.Run("save exist", func(t *testing.T) {
		//连接数据库
		db := prepareDB(t)
		defer db.Close()
		svc := &mysqlStudentService{
			db,
		}
		//创表
		svc.createStudentTable()
		//test结束后删表
		defer svc.dropTable()

		s := &model.Student{
			Id:   1,
			Name: "weidongqi",
			Age:  22,
		}
		svc.SaveStudent(s)
		wantErr := true
		if err := svc.SaveStudent(s); (err != nil) != wantErr {
			t.Errorf("saveStudent() error = %v, wantErr %v", err, wantErr)
		}
	})

}

func TestMysqlStudentService_UpdateStudent(t *testing.T) {

	t.Run("update not exist", func(t *testing.T) {
		//连接数据库
		db := prepareDB(t)
		svc := &mysqlStudentService{
			db,
		}
		svc.createStudentTable()
		//test结束后删表
		defer svc.dropTable()

		s := &model.Student{
			Id:   8,
			Name: "weidongqi",
			Age:  22,
		}
		wantErr := true
		if err := svc.UpdateStudent(s); (err != nil) != wantErr {
			t.Errorf("saveStudent() error = %v, wantErr %v", err, wantErr)
		}

	})
	t.Run("update exist", func(t *testing.T) {
		//连接数据库
		db := prepareDB(t)

		//defer db.Close()
		svc := &mysqlStudentService{
			db,
		}
		//创表
		svc.createStudentTable()
		//test结束后删表
		defer svc.dropTable()
		s1 := &model.Student{
			Id:   1,
			Name: "weidongqi",
			Age:  22,
		}
		svc.SaveStudent(s1)

		s := &model.Student{
			Id:   1,
			Name: "xiaoxing",
			Age:  33,
		}
		wantErr := false
		if err := svc.UpdateStudent(s); (err != nil) != wantErr {
			t.Errorf("saveStudent() error = %v, wantErr %v", err, wantErr)
		}
		s2, err2 := svc.GetStudent(1)
		if err2 != nil {
			t.Error(err2)
		}
		if !reflect.DeepEqual(s2, s) {
			t.Errorf("getStudentById() got = %v, want %v", s2, s)
		}
	})

}
