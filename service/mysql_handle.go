package service

import (
	"dc/driver"
	"dc/err"
	"dc/util"
	"fmt"
)

//定义服务
type MysqlHandleService struct {
}

var queryBuilder  *QueryBuilderS
//启动参数处理
func initQuery(args Args)  {
	queryBuilder = &QueryBuilderS{}
	InitArgs(args)
}
//服务里的方法
//查询单条或多条, return 用[]map
func (mhs MysqlHandleService) Fetch(args Args) {
	initQuery(args)

	if Ret.Code != 0{
		return
	}

	InitField(queryBuilder)

	InitWhere(queryBuilder)

	if Ret.Code != 0{
		return
	}
	InitOrder(queryBuilder)

	if Ret.Code != 0{
		return
	}
	InitJoin(queryBuilder)
	if Ret.Code != 0{
		return
	}
	MysqlEngine := driver.GetMysqlEngine()

	engine := MysqlEngine.Table(ArgsBuilder.Table).
		Where(queryBuilder.Query, queryBuilder.BindValue...).
		Limit(ArgsBuilder.Limit, ArgsBuilder.Offset)

	//别名
	if ArgsBuilder.Alias != ""{
		engine = engine.Alias(ArgsBuilder.Alias)
	}
	if queryBuilder.Join != nil{
		for _,val := range queryBuilder.Join{
			engine = engine.Join(val[0].(string), val[1], val[2].(string))
		}
	}
	if queryBuilder.Field != nil && ArgsBuilder.Omit != nil{
		Ret.Code = err.INVALID_PARAMS
		return
	}
	//字段
	if queryBuilder.Field != nil{
		engine = engine.Cols(queryBuilder.Field...)
	}
	//字段
	if ArgsBuilder.Omit != nil{
		engine = engine.Omit(ArgsBuilder.Omit...)
	}
	if ArgsBuilder.Distinct != nil{
		engine = engine.Distinct(ArgsBuilder.Distinct...)
	}
	if ArgsBuilder.Having != ""{
		engine = engine.Having(ArgsBuilder.Having)
	}
	//排序
	if queryBuilder.Order != nil{
		for k,val := range queryBuilder.Order {
			if len(val)<1{
				continue
			}
			if k == "desc" {
				engine = engine.Desc(val...)
			}
			if k == "asc"{
				engine = engine.Asc(val...)
			}
		}
	}

	var data []map[string]interface{}

	error := engine.Find(&data)

	if error != nil{
		Error(error)
		return
	}

	//转换格式 特别对 []byte 进行处理
	Ret.Data = FormatValue(data)
}

//查询是否存在
func (mhs MysqlHandleService) Exist(args Args) {

	initQuery(args)

	if Ret.Code != 0{
		return
	}

	InitWhere(queryBuilder)

	if Ret.Code != 0{
		return
	}

	MysqlEngine := driver.GetMysqlEngine()

	has, error := MysqlEngine.Table(ArgsBuilder.Table).
		Where(queryBuilder.Query, queryBuilder.BindValue...).Exist()

	if error != nil{
		Error(error)
		return
	}

	Ret.Data = has
}
//创建数据(单条或多条)
func (mhs MysqlHandleService) Create(args Args){
	initQuery(args)

	if ArgsBuilder.Body == nil{
		Ret.Code = err.INVALID_PARAMS
		return
	}
	var tmp  []Args
	for _, val := range ArgsBuilder.Body.([]interface{}) {
		tmp = append(tmp, val.(map[string]interface{}))
	}
	fun := TableBindToModel[ArgsBuilder.Table]

	var  multiModels  []interface{}

	for _,arg := range tmp {
		model := fun(arg)
		multiModels = append(multiModels, model)
	}
	MysqlEngine := driver.GetMysqlEngine()

	affected, error := MysqlEngine.Insert(multiModels...)

	if error != nil{
		Error(error)
		return
	}

	if affected == 0{
		Ret.Code = err.ERROR_INSERT_FAIL
		return
	}

	funResolver := TableResolverToModel[ArgsBuilder.Table]
	funResolver(multiModels, len(args))
}

//修改
func (mhs MysqlHandleService) Update(args Args){
	initQuery(args)

	InitWhere(queryBuilder)

	if ArgsBuilder.Body == nil{
		Ret.Code = err.INVALID_PARAMS
		return
	}
	//参数绑定
	fun := TableBindToModel[ArgsBuilder.Table]

	model := fun(ArgsBuilder.Body.(Args))

	MysqlEngine := driver.GetMysqlEngine()

	_, error := MysqlEngine.Where(queryBuilder.Query, queryBuilder.BindValue...).Update(model)

	if error != nil{
		Error(error)
		return
	}
}
//根据条件删除
func (mhs MysqlHandleService) Delete(args Args)  {
	initQuery(args)

	InitWhere(queryBuilder)

	fun := TableBindToModel[ArgsBuilder.Table]

	model := fun(nil)
	MysqlEngine := driver.GetMysqlEngine()

	_, error := MysqlEngine.Where(queryBuilder.Query, queryBuilder.BindValue...).Delete(model)

	if error != nil{
		Error(error)
		return
	}
}
//查询一条
func (mhs MysqlHandleService) Get(args Args){
	initQuery(args)
	InitField(queryBuilder)
	if ArgsBuilder.Id == 0{
		Ret.Code = err.INVALID_PARAMS
		return
	}
	var data []map[string]interface{}
	MysqlEngine := driver.GetMysqlEngine()

	engine := MysqlEngine.Table(ArgsBuilder.Table).ID(ArgsBuilder.Id)

	//字段
	if queryBuilder.Field != nil{
		engine = engine.Cols(queryBuilder.Field...)
	}

	bool, error := engine.Get(&data)
	if !bool{
		Ret.Code = err.ERROR_NOT_EXIST
	}
	if error != nil{
		Error(error)
		return
	}
	Ret.Data = data
}
//修改单条
func (mhs MysqlHandleService) Set(args Args) {
	initQuery(args)
	if ArgsBuilder.Id == 0{
		Ret.Code = err.INVALID_PARAMS
		return
	}
	if ArgsBuilder.Body == nil{
		Ret.Code = err.INVALID_PARAMS
		return
	}
	fun := TableBindToModel[ArgsBuilder.Table]

	model := fun(ArgsBuilder.Body.(map[string]interface{}))

	MysqlEngine := driver.GetMysqlEngine()

	_, error := MysqlEngine.ID(ArgsBuilder.Id).Update(&model)

	if error != nil{
		Error(error)
		return
	}

}
//删除单条
func (mhs MysqlHandleService) Del(args Args) {
	initQuery(args)
	if ArgsBuilder.Id == 0{
		Ret.Code = err.INVALID_PARAMS
		return
	}
	fun := TableBindToModel[ArgsBuilder.Table]

	model := fun(map[string]interface{}{})
	MysqlEngine := driver.GetMysqlEngine()

	_, error := MysqlEngine.Table(ArgsBuilder.Table).ID(ArgsBuilder.Id).Delete(&model)
	if error != nil{
		Error(error)
		return
	}
}
//原生sql语句查询
func (mhs MysqlHandleService)Query(args Args)  {
	initQuery(args)
	query, ok:= args["query"]
	if !ok{
		Ret.Code = err.INVALID_PARAMS
		return
	}
	Qargs, ok:= args["args"]
	if !ok{
		Ret.Code = err.INVALID_PARAMS
		return
	}
	MysqlEngine := driver.GetMysqlEngine()
	engine := MysqlEngine.Table(ArgsBuilder.Table).SQL(query, Qargs.([]interface{})...)
	var data []map[string]interface{}

	error := engine.Find(&data)

	if error != nil{
		Error(error)
		return
	}

	Ret.Data = FormatValue(data)
}
//sum 统计字段的和
func (mhs MysqlHandleService) Sum(args Args) {

	initQuery(args)

	if Ret.Code != 0{
		return
	}

	InitWhere(queryBuilder)

	if Ret.Code != 0{
		return
	}
	sum, ok:= args["sum"]
	if !ok{
		Ret.Code = err.INVALID_PARAMS
		return
	}
	fun := TableBindToModel[ArgsBuilder.Table]

	model := fun(ArgsBuilder.Body.(map[string]interface{}))
	MysqlEngine := driver.GetMysqlEngine()

	has, error := MysqlEngine.Table(ArgsBuilder.Table).
		Where(queryBuilder.Query, queryBuilder.BindValue...).Sums(model, sum.([]string)...)

	if error != nil{
		Error(error)
		return
	}

	Ret.Data = has
}
//开启事务,,用AffairHandle 转入事务句柄,返回前端一个session 的字符串,作为键值
//在执行操作的时候,需要把session 键值传输过来,
//根据是否有有session键值,,去AffairHandle 里面找到真正的句柄
//根据句柄执行操作
//当遇到回滚或提交或程序中断的时候,,,清除AffairHandle里面的句柄
func test()  {
	MysqlEngine := driver.GetMysqlEngine()
	session := MysqlEngine.NewSession()
	defer session.Close()
}
//转换格式 特别对 []byte 进行处理
func FormatValue(dataMap []map[string]interface{})[]map[string]interface{}{

	if length := len(dataMap);length>0{
		for i:=0;i<length ;i++  {
			util.FormatValue(&dataMap[i])
		}
	}
	return  dataMap
}
//错误处理
func Error(error error){
	Ret.Code = err.ERROR
	Ret.Msg = fmt.Sprintf(error.Error())
}
