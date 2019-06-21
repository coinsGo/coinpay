package config

import (
	"github.com/go-ini/ini"
	"log"
)

type global struct {
	Debug          	bool
	ServerPort   	string
	ServerWebsite 	string
	LogPath      	string

	DbLogMode         bool    //数据库日志模式，开启true, 关闭false
	DbMaxIdleConns    int    //最大空闲连接数
	DbMaxOpenConns    int   //最大连接数
	DbConnMaxLifetime int   //mysql超时时间
}

type envdata struct {
	MysqlIp       string
	MysqlPort     string
	MysqlUsername string
	MysqlPassword string
	MysqlDbname   string
	MysqlPrefix   string
}

var Global = &global{}
var Envdata = &envdata{}


func Setup(envname string) {

	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Fatalf("config.Setup, fail to parse './config/config.ini': %v", err)
	}
	mapTo(cfg,"GLOBAL", Global)
	mapTo(cfg, envname, Envdata)
}


func mapTo( cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
