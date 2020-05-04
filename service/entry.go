//中间件
//主要功能是进行数据转换,根据相关驱动,进行不同方法调用

package service

import (
	"dc/err"
	"dc/util"
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/hprose/hprose-golang/rpc"
	"reflect"
)


//入口中间件 处理区分驱动,进行
func EntryInvokeHandler(
	name string,
	args []reflect.Value,
	context rpc.Context,
	next rpc.NextInvokeHandler) (results []reflect.Value, err error) {

	ret := dispatcher(name, args[0].String())

	results = []reflect.Value{
		reflect.ValueOf(ret),
	}
	//results, err = next(name, args, context)
	return
}

//对于用户来说 mysql mongodb redis,没什么区别,都是统一的方法,不同的参数
//而对于数据中心来说,需要抽象封装一层,让前端用户感觉不到,到底是什么驱动
//数据中心,在接到方法调用后,需要根据驱动不同,去调用不同的驱动方法,每个驱动都有自己的一套解析器,查询方式等等
//那么就要借用中间件,在程序入口处,统一处理,
func dispatcher(name, args string) []byte {
	Ret = ResBody{Code: err.SUCCESS}
	JsonArgs := initArgs(args)
	driver   := getDriver(JsonArgs)
	if Ret.Code != 0{
		return Resp(Ret)
	}
	handle   := getHandle(driver)
	if Ret.Code != 0{
		return Resp(Ret)
	}
	switch name {
	case "DC_Fetch":
		handle.Fetch(JsonArgs)
	case "DC_Update":
		handle.Update(JsonArgs)
	case "DC_Delete":
		handle.Delete(JsonArgs)
	case "DC_Create":
		handle.Create(JsonArgs)
	case "DC_Exist":
		handle.Exist(JsonArgs)
	case "DC_Count":
		//ret = handle.Count(JsonArgs)
	case "DC_Get":
		handle.Get(JsonArgs)
	case "DC_Set":
		handle.Set(JsonArgs)
	case "DC_Del":
		handle.Del(JsonArgs)
	case "DC_Query":
		handle.Query(JsonArgs)
	default:
	// todo 方法错误,没有该方法
		Ret.Code = err.ERROR_NOT_FUN
	}
	return  Resp(Ret)
}
//得到驱动操作服务
func getHandle(driver string) HandleInterface {
	var handle HandleInterface
	switch driver {
	case "mysql":
		handle = MysqlHandleService{}
	default:
		Ret.Code = err.ERROR_NOT_DRIVER
	}
	return handle
}

//解析参数 json->map
func initArgs(args string) Args {
	jsonEntity ,_ := simplejson.NewJson([]byte(args))
	JsonArgs,_ := jsonEntity.Map()
	return JsonArgs
}
//分析当前调用的驱动, default mysql
func getDriver(JsonArgs Args)string{
	driver := "mysql"

	//驱动处理
	driverName, ok:= JsonArgs["driver"]
	if ok && driverName != ""{
		driver = driverName.(string)
	}
	if util.ContainsString(util.Cfg.Driver.DriverName, driver)<0{
		//todo 不支持当前驱动 ,支持:mysql,mongodb,redis
		Ret.Code = err.ERROR_NOT_DRIVER
	}
	return driver
}
//返回函数
func Resp(resBody ResBody) []byte {
	//设置header 因为默认是text/html ,现在设置为application/json

	if resBody.Msg == ""{
		resBody.Msg = err.Msg(resBody.Code)
	}
	if resBody.Code !=0{
		resBody.Data = nil
	}

	ret,_:=  json.Marshal(resBody)
	return ret
}
