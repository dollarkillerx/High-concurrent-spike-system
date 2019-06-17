/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 下午8:58
* */
package mysql_conn

import (
	"High-concurrent-spike-system/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	MySQL *sql.DB
	err   error
)

func init() {
	MySQL, err = sql.Open("mysql", config.Config.MySQLDsn)
	if err != nil {
		panic(err.Error())
	}
}
