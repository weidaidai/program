# Database  program
> 用golang 连接 mysql 和redis

第一种方法 -在database/sql硬编码

go连接mysql 使用标准库为

> database/sql
>
> 使用的methods

```go
//open()校验参数是否正确
func Open(driverName string, dataSourceName string) (*DB, error)
//driverName 数据库驱动名mysql,dataSourceName 标准为user:password@tcp(url:port)/database_name



//ping()连接数据库
func (db *DB) Ping() error

//查询单行
func (db *DB) QueryRow(query string, args ...interface{}) *Row

//查询多行
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
//参数args表示query中的占位参数,?里面的补位

//
```

> github.com/go-sql-driver/mysql

DB是一个数据库句柄，表示0个或多个底层连接池。它对于多个goroutine并发使用是安全的，里面有很多methods可以调用

一般都会初始化一个全局对象来使用

> var db *sql.DB

```go
SetMaxOpenConns() // 与数据库连接的最大连接数

SetMaxIdleConns() // 与连接池连接的最大空闲连接数
```

单行查询

