package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Jordanzuo/GameServer_Go/src/config"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"time"
)

var (
	gameModelDB *sql.DB
	gameDB      *sql.DB
	redisPool   *redis.Pool
)

// 初始化数据库连接相关的配置
func init() {
	gameModelDB = openMysqlConnection(config.GameModelDBConnection)
	gameDB = openMysqlConnection(config.GameDBConnection)
	redisPool = newRedisPool(config.RedisConnection)
}

// 获取连接游戏数据库的DB对象
func GameDB() *sql.DB {
	return gameDB
}

// 获取连接游戏模型数据库的DB对象
func GameModelDB() *sql.DB {
	return gameModelDB
}

// 获取Redis的连接
// 返回值
// Redis连接对象
func RedisConn() redis.Conn {
	return redisPool.Get()
}

func openMysqlConnection(connectionString string) *sql.DB {
	connectionSlice := strings.Split(connectionString, ";")

	// 建立数据库连接
	db, err := sql.Open("mysql", connectionSlice[0])
	if err != nil {
		panic(errors.New(fmt.Sprintf("打开游戏数据库失败,连接字符串为：%s", connectionString)))
	}

	// 设置连接池相关
	maxOpenConns_string := strings.Replace(connectionSlice[1], "MaxOpenConns=", "", 1)
	maxOpenCons, err := strconv.Atoi(maxOpenConns_string)
	if err != nil {
		panic(errors.New(fmt.Sprintf("MaxOpenConns必须为int型,连接字符串为：%s", connectionString)))
	}

	maxIdleConns_string := strings.Replace(connectionSlice[2], "MaxIdleConns=", "", 1)
	maxIdleConns, err := strconv.Atoi(maxIdleConns_string)
	if err != nil {
		panic(errors.New(fmt.Sprintf("MaxIdleConns必须为int型,连接字符串为：%s", connectionString)))
	}

	db.SetMaxOpenConns(maxOpenCons)
	db.SetMaxIdleConns(maxIdleConns)
	db.Ping()

	return db
}

func newRedisPool(connectionString string) *redis.Pool {
	connectionSlice := strings.Split(connectionString, ";")

	// 获取连接池相关
	maxActive_string := strings.Replace(connectionSlice[1], "MaxActive=", "", 1)
	maxActive, err := strconv.Atoi(maxActive_string)
	if err != nil {
		panic(errors.New(fmt.Sprintf("MaxActive必须为int型,连接字符串为：%s", connectionString)))
	}

	maxIdle_string := strings.Replace(connectionSlice[2], "MaxIdle=", "", 1)
	maxIdle, err := strconv.Atoi(maxIdle_string)
	if err != nil {
		panic(errors.New(fmt.Sprintf("MaxIdle必须为int型,连接字符串为：%s", connectionString)))
	}

	idleTimeout_string := strings.Replace(connectionSlice[3], "IdleTimeout=", "", 1)
	idleTimeout_int, err := strconv.Atoi(idleTimeout_string)
	if err != nil {
		panic(errors.New(fmt.Sprintf("IdleTimeout必须为int型,连接字符串为：%s", connectionString)))
	}
	idleTimeout := time.Duration(idleTimeout_int) * time.Second

	password := strings.Replace(connectionSlice[4], "Password=", "", 1)

	return &redis.Pool{
		MaxIdle:     maxIdle,     // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
		MaxActive:   maxActive,   // 最大的激活连接数，表示同时最多有N个连接
		IdleTimeout: idleTimeout, //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", connectionSlice[0])
			if err != nil {
				return nil, err
			}

			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// 处理Redis错误
// err：错误对象
func HandleRedisError(err error) {
	if err != nil {
		panic(errors.New(fmt.Sprintf("执行Redis命令失败，错误信息为：%s", err)))
	}
}

// 处理Mysql错误
// err：错误对象
func HandleMysqlError(err error) {
	if err != nil {
		panic(errors.New(fmt.Sprintf("执行Mysql命令失败，错误信息为：%s", err)))
	}
}

// 处理其它错误
// err：错误对象
func HandleOtherError(err error) {
	if err != nil {
		panic(errors.New(fmt.Sprintf("其他错误，错误信息为：%s", err)))
	}
}
