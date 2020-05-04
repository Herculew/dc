package util

import (
	"github.com/go-ini/ini"
	"time"
)

type AppCfg struct{
	PageSize int `ini:"PAGE_SIZE"`
	ServerName string `ini:"SERVER_NAME"`
}

type MongodbCfg struct{
	MongodbUri string `ini:"MONGODB_URI"`
	TimeOut time.Duration `ini:"TIMEOUT"`
}

type RpcServerCfg struct{
	Host string `ini:"HOST"`
	Port string   `ini:"PORT"`
}

type HttpServerCfg struct{
	Host string `ini:"HOST"`
	Port int32   `ini:"PORT"`
	ReadTime time.Duration   `ini:"READ_TIMEOUT"`
	WriteTime time.Duration   `ini:"WRITE_TIMEOUT"`
}

type MysqlCfg struct{
	Host string `ini:"HOST"`
	TablePrefix string `ini:"TABLE_PREFIX"`
	User string `ini:"USER"`
	Password string `ini:"PASSWORD"`
	Database string `ini:"DATABASE"`
}

//type RedisCfg struct{
//	Host string
//	Port int32
//	Auth string
//}
type DriverCfg struct{
	DriverName []string `ini:"DRIVER_NAME"`
}

type Config struct {
	RunMode string `ini:"RUN_MODE"`
	ServerName string `ini:"SERVER_NAME"`
	Mysql MysqlCfg
	App AppCfg
	Mongodb MongodbCfg
	HttpServer HttpServerCfg
	RpcServer RpcServerCfg
	Driver DriverCfg
}

//全局变量
var Cfg = &Config{}

//加载配置
func InitConfig(ConfigFile string) error {
	return ini.MapTo(Cfg, ConfigFile)
}