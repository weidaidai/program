package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"testing"
)

func prepareDB(t *testing.T) *sql.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	db, err := openDB(dsn)
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

func insertTestStudent(t *testing.T, db *sql.DB, s *Student) {
	err := saveStudent(db, s)
	if err != nil {
		t.Error(err)
	}
}

func Test_openDB(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		wantDb  bool
		wantErr bool
	}{
		{
			name:    "good case",
			args:    args{dsn: "root:123456@tcp(127.0.0.1:3306)/sql_domo"},
			wantDb:  true,
			wantErr: false,
		},
		{
			name:    "bad case",
			args:    args{dsn: "root:1234567@tcp(127.0.0.1:3306)/sql_domo"},
			wantDb:  false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDb, err := openDB(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("openDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDb != nil {
				gotDb.Close()
			}

		})
	}
}

// TODO 把STUDENT表删掉
func Test_dropTable(t *testing.T) {
	db := prepareDB(t)
	prepareTable(t, db)
	defer db.Close()

	t.Run("", func(t *testing.T) {

		wantErr := false

		if err := dropTable(db); (err != nil) != wantErr {
			t.Errorf("dropTable() error = %v, wantErr %v", err, wantErr)
		}
		_, err2 := listAllStudents(db)
		if err2 == nil {
			t.Error(err2)
		}
	})

}

// TODO 创STUDENT表
func Test_createStudentTable(t *testing.T) {

	db := prepareDB(t)
	defer db.Close()

	t.Run("create table", func(t *testing.T) {
		defer dropTable(db)

		wantErr := false
		if err := createStudentTable(db); (err != nil) != wantErr {
			t.Errorf("createStudentTable() error = %v, wantErr %v", err, wantErr)
		}

		_, err2 := listAllStudents(db)
		if err2 != nil {
			t.Error(err2)
		}
	})

}

func Test_saveStudent(t *testing.T) {

	//连接数据库
	db := prepareDB(t)

	defer db.Close()

	//test结束后删表
	defer dropTable(db)

	t.Run("save not exist", func(t *testing.T) {
		prepareTable(t, db)
		defer dropTable(db)
		s := &Student{
			Id:   1,
			Name: "weidongqi",
			Age:  22,
		}

		wantErr := false
		if err := saveStudent(db, s); (err != nil) != wantErr {
			t.Errorf("saveStudent() error = %v, wantErr %v", err, wantErr)
		}
		s2, err2 := getStudentById(db, 1)
		if err2 != nil {
			t.Error(err2)
		}
		if !reflect.DeepEqual(s2, s) {
			t.Errorf("getStudentById() got = %v, want %v", s2, s)
		}
	})

	t.Run("save exist", func(t *testing.T) {
		prepareTable(t, db)
		defer dropTable(db)

		s := &Student{
			Id:   1,
			Name: "weidongqi",
			Age:  22,
		}
		insertTestStudent(t, db, s)
		wantErr := true
		if err := saveStudent(db, s); (err != nil) != wantErr {
			t.Errorf("saveStudent() error = %v, wantErr %v", err, wantErr)
		}
	})
}

func Test_updateStudent(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	t.Run("update exist", func(t *testing.T) {
		prepareTable(t, db)
		defer dropTable(db)

		s1 := &Student{Id: 1, Name: "weidongqi", Age: 22}
		insertTestStudent(t, db, s1)

		s2 := &Student{
			Id:   1,
			Name: "weidongqi2",
			Age:  33,
		}
		wantErr := false
		if err := updateStudent(db, s2); (err != nil) != wantErr {
			t.Errorf("updateStudent() error = %v, wantErr %v", err, wantErr)
		}
		s3, err2 := getStudentById(db, 1)
		if err2 != nil {
			t.Error(err2)
		}
		if !reflect.DeepEqual(s3, s2) {
			t.Errorf("getStudentById() got = %v, want %v", s3, s2)
		}
	})

	t.Run("update not exist", func(t *testing.T) {
		prepareTable(t, db)
		defer dropTable(db)

		s2 := &Student{
			Id:   2,
			Name: "weidongqi2",
			Age:  33,
		}
		wantErr := false
		if err := updateStudent(db, s2); (err != nil) != wantErr {
			t.Errorf("updateStudent() error = %v, wantErr %v", err, wantErr)
		}
	})

}

func Test_deleteStudent(t *testing.T) {
	db := prepareDB(t)
	defer db.Close()

	t.Run("delete exist", func(t *testing.T) {
		prepareTable(t, db)
		defer dropTable(db)

		s1 := &Student{Id: 1, Name: "weidongqi", Age: 22}
		insertTestStudent(t, db, s1)

		wantErr := false
		if err := deleteStudent(db, 1); (err != nil) != wantErr {
			t.Errorf("deleteStudent() error = %v, wantErr %v", err, wantErr)
		}
		s2, err2 := getStudentById(db, 1)
		if err2 != nil {
			t.Error(err2)
		}
		if s2 == nil {
			t.Errorf("getStudentById() got = %v, want %v", s2, nil)
		}
	})

	t.Run("delete not exist", func(t *testing.T) {
		prepareTable(t, db)
		defer dropTable(db)
		wantErr := false
		if err := deleteStudent(db, 2); (err != nil) != wantErr {
			t.Errorf("deleteStudent() error = %v, wantErr %v", err, wantErr)
		}
	})

}

func Test_listAllStudents(t *testing.T) {
	//连接数据库
	db := prepareDB(t)

	defer db.Close()
	//创表
	prepareTable(t, db)
	//test结束后删表
	defer dropTable(db)
	//插入数据
	s1 := &Student{Id: 1, Name: "weidongqi", Age: 22}
	s2 := &Student{Id: 2, Name: "weidongqi2", Age: 33}
	insertTestStudent(t, db, s1)
	insertTestStudent(t, db, s2)

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    []*Student
		wantErr bool
	}{
		{
			name: "",
			args: args{db: db},
			want: []*Student{
				{Id: 1, Name: "weidongqi", Age: 22},
				{Id: 2, Name: "weidongqi2", Age: 33},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := listAllStudents(tt.args.db)
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
