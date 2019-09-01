package pgsql_conn

import (
	"High-concurrent-spike-system/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	"log"
	"time"
)

var (
	PgDb *gorm.DB
	err error
)

func init() {
	PgDb, err = gorm.Open("postgres", config.MyConfig.Pgsql.Dsn)
	if err != nil {
		panic(err.Error())
	}

	ping := PgDb.DB().Ping()
	if ping != nil {
		log.Println("DB Error=================================")
		log.Println("数据库链接失败 Failure of database link")
		log.Println("DB Error=================================")
	}

	PgDb.DB().SetConnMaxLifetime(time.Hour * config.MyConfig.Pgsql.TimeOut)
	PgDb.DB().SetMaxIdleConns(config.MyConfig.Pgsql.MaxIdle)
	PgDb.DB().SetMaxOpenConns(config.MyConfig.Pgsql.MaxOpen)

	mapping()

}

// 数据库映射
func mapping() {

}
