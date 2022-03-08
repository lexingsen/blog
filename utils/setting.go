package utils

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	ServerMode string
	ServerPort string
	DBType     string
	DBhost     string
	DBuser     string
	DBpassword string
	DBdatabase string
	DBport     string
)

func init() {

	file, err := ini.Load("config/config.ini")
	if err != nil {
		log.Error("load config fail")
	}
	loadMySQLConfig(file)
	loadServerConfig(file)
}

func loadServerConfig(file *ini.File) {
	ServerMode = file.Section("server").Key("ServerMode").MustString("debug")
	ServerPort = file.Section("server").Key("ServerPort").MustString(":8000")
}

func loadMySQLConfig(file *ini.File) {
	DBType = file.Section("database").Key("DBType").MustString("mysql")
	DBhost = file.Section("database").Key("DBhost").MustString("localhost")
	DBuser = file.Section("database").Key("DBuser").MustString("root")
	DBpassword = file.Section("database").Key("DBpassword").MustString("111111")
	DBdatabase = file.Section("database").Key("DBdatabase").MustString("blog")
	DBport = file.Section("database").Key("DBport").MustString("3306")
}
