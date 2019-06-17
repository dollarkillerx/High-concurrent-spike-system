/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-17
* Time: 下午8:59
* */
package config

import (
	"encoding/json"
	"os"
)

type Conf struct {
	Host     string `json:"host"`
	Debug    bool   `json:"debug"`
	MySQLDsn string `json:"mysql_dsn"`
}

var (
	Config *Conf
)

func init() {
	path := "./config.json"
	file, e := os.Open(path)
	if e != nil {
		panic(e.Error())
	}

	decoder := json.NewDecoder(file)

	Config = &Conf{}
	e = decoder.Decode(Config)
	if e != nil {
		panic(e.Error())
	}
}
