package config

import (
	"errors"
	"log"
	"path/filepath"
	"runtime"

	"github.com/go-ini/ini"
)

type global struct {
	Debug         bool
	ServerPort    string
	ServerWebsite string
	LogPath       string

	DbLogMode         bool //数据库日志模式，开启true, 关闭false
	DbMaxIdleConns    int  //最大空闲连接数
	DbMaxOpenConns    int  //最大连接数
	DbConnMaxLifetime int  //mysql超时时间
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

func Setup(env string) {

	file := currentDir() + "/config.ini"
	cfg, err := ini.Load(file)
	if err != nil {
		log.Fatalf("config.Setup, fail to parse '"+file+"': %v", err)
	}
	mapTo(cfg, "GLOBAL", Global)
	mapTo(cfg, env, Envdata)
}

func mapTo(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}

func currentDir() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	return filepath.Dir(file)
}
