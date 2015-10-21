package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Jordanzuo/GameServer_Go/src/config"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

var (
	gameModelDB *sql.DB
	gameDB      *sql.DB
)

// 初始化数据库连接相关的配置
func init() {
	gameModelDB = openConnections(config.GameModelDBConnection)
	gameDB = openConnections(config.GameDBConnection)
}

// 获取连接游戏数据库的DB对象
func GameDB() *sql.DB {
	return gameDB
}

// 获取连接游戏模型数据库的DB对象
func GameModelDB() *sql.DB {
	return gameModelDB
}

func openConnections(connectionString string) *sql.DB {
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

	return db
}
