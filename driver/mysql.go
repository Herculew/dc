package driver

import (
	"dc/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"time"
	"xorm.io/core"
)

var MysqlEngine *xorm.Engine

const dbType  = "mysql"


type MysqlMapper struct {
	Query string     //查询的条件
	Field []string     //字段
	BindValue []interface{}
	Order map[string][]string
	engine *xorm.Engine
}


func GetMysqlEngine()*xorm.Engine{

	if MysqlEngine != nil{
		return MysqlEngine
	}

	DsName  := util.Cfg.Mysql.User +":"+util.Cfg.Mysql.Password+"@"+util.Cfg.Mysql.Host+"/"+util.Cfg.Mysql.Database+"?charset=utf8"

	mysqlEngine ,err := xorm.NewEngine(dbType,DsName)
	if err != nil{
		log.Fatal(err.Error())
	}

	// 创建引擎组
	// https://github.com/go-xorm/xorm/blob/master/README_CN.md


	//如果需要设置连接池的空闲数大小，可以使用engine.SetMaxIdleConns()来实现。
	//如果需要设置最大打开连接数，则可以使用engine.SetMaxOpenConns()来实现。

	mysqlEngine.ShowSQL(true) //是否显示sql语句
	mysqlEngine.SetMaxOpenConns(10) //设置数据库最大的连接数
	//自动

	//tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, util.Cfg.Mysql.TablePrefix)
	//mysqlEngine.SetTableMapper(tbMapper)

	//设置时钟
	mysqlEngine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	err1 := mysqlEngine.Ping()
	if err1 != nil{
		log.Fatal(err1.Error())
	}

	MysqlEngine = mysqlEngine
	return MysqlEngine
}
func (mm *MysqlMapper)GetMysqlEngine1()*MysqlMapper{

	if MysqlEngine != nil{
		mm.engine = MysqlEngine
		return mm
	}

	DsName  := util.Cfg.Mysql.User +":"+util.Cfg.Mysql.Password+"@"+util.Cfg.Mysql.Host+"/"+util.Cfg.Mysql.Database+"?charset=utf8"

	mysqlEngine ,err := xorm.NewEngine(dbType,DsName)
	if err != nil{
		log.Fatal(err.Error())
	}

	// 创建引擎组
	// https://github.com/go-xorm/xorm/blob/master/README_CN.md


	//如果需要设置连接池的空闲数大小，可以使用engine.SetMaxIdleConns()来实现。
	//如果需要设置最大打开连接数，则可以使用engine.SetMaxOpenConns()来实现。

	mysqlEngine.ShowSQL(true) //是否显示sql语句
	mysqlEngine.SetMaxOpenConns(10) //设置数据库最大的连接数
	//自动

	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, util.Cfg.Mysql.TablePrefix)
	mysqlEngine.SetTableMapper(tbMapper)

	//设置时钟
	mysqlEngine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	err1 := mysqlEngine.Ping()
	if err1 != nil{
		log.Fatal(err1.Error())
	}

	mm.engine = mysqlEngine
	return mm
}
func (mm *MysqlMapper)Where()  {

}
func CloseDB() {
	defer MysqlEngine.Close()
}