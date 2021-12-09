# Database  program

### MySQL

> 用golang 连接 mysql 和redis

第一种方法 -在database/sql硬编码

go连接mysql 使用标准库为

> database/sql
>
> 使用的methods

```go
_ "github.com/go-sql-driver/mysql"
```

> DB是一个数据库句柄，表示0个或多个底层连接池。它对于多个goroutine并发使用是安全的，里面有很多methods可以调用
>
> 一般都会初始化一个全局对象来使用
>
> var db *sql.DB

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

//更新，删除，插入


```

> 

```go
SetMaxOpenConns() // 与数据库连接的最大连接数

SetMaxIdleConns() // 与连接池连接的最大空闲连接数
```

### Redis

```go
"github.com/go-redis/redis" 
```

client是一个 Redis 客户端，代表零个或多个底层连接池。多个 goroutine 并发使用是安全的。

里面有很多methods可以调用

一般都会初始化一个全局对象来使用

> var rdb  *redis.Client

```go
//取options类型对象的地址常用 addr password DB, poolsize
/*
源码包
type Options struct {
    Network            string
    Addr               string
    Dialer             func() (net.Conn, error)
    OnConnect          func(*Conn) error
    Password           string
    DB                 int
    MaxRetries         int
    ....*/
    
   // NewClient 返回一个客户端给Options指定的Redis Server
    //func NewClient(opt *Options) *Client
    // 例如 
    rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,//一般有6个
		PoolSize: 100,
	})
   //连接 
   // func (c *cmdable) Ping() *StatusCmd
_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}


/*    
源码包
type StatusCmd struct {
    baseCmd
    val string
}
Methods:
Val() string
Result() (string, error) //一般使用返回结果
String() string
readReply(rd *proto.Reader) error

```

string 使用

```go

func (c *cmdable) Set(key string, value interface{}, expiration time.Duration) *StatusCmd

```

zset 使用

```go
/* 
源码包
type Z struct {
    Score  float64
    Member interface{}
}
*/
添加多个数据
zsetkey := "database"
	languages := []redis.Z{
		redis.Z{95, "mysql"},
		redis.Z{66, "redis1"},
		redis.Z{77, "mogondb1"},
	}
num, err := rdb.ZAdd(zsetkey, languages...).Result()
。。。。

```

### Http

![](image\http.png)

服务器在接收到请求时，首先会进入路由(`router`)，也成为服务复用器（`Multiplexe`），路由的工作在于请求找到对应的处理器(`handler`)，处理器对接收到的请求进行相应处理后构建响应并返回给客户端。Go实现的`http server`同样遵循这样的处理流程。

go 启动http server 一般常见有两种

我一般使用HandleFunc

