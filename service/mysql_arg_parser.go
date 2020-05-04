//mysql参数解析器
package service

import (
	"dc/err"
	"dc/util"
	"fmt"
	"strings"
)
//定义常用的数据类型
var (
	//常用操作符
	WhereOption = []string{"=", "in", "not in", ">", ">=", "<", "<=", "like","field_in_set", "not between","between" }

	//排序操作符
	OrderOption = []string{"desc", "asc"}
	//链接操作符
	LinkOption = []string{ "or", "and"}

	ArgsBuilder ArgsBuilderS

)

//定义参数结构体
type ArgsBuilderS struct {
	Driver string  // "mysql"
	Table string
	Alias string     //别名
	Where interface{}
	Field  string  // "*"
	Having  string  // "*"
	Limit  int     //1
	Order interface{}
	Join []interface{} //联合
	Omit []string
	Distinct []string
	Group []interface{}
	Body interface{}
	Offset int
	Id int
}


//where 结构体
type QueryBuilderS struct {
	Query string     //查询的条件
	Field []string     //字段
	BindValue []interface{}
	Order map[string][]string
	Join [][]interface{}
}

// 参数处理
// todo  还要验证 数据类型是否正确
// 因为很多时候不是数组 之类的 还有可以在php端进行限制

func InitArgs(JsonArgs Args){

	ArgsBuilder = ArgsBuilderS{
		Limit:1,
	}
	//表名处理
	if table, ok:= JsonArgs["table"]; ok {
		if table != ""{
			name, alias := util.Split2Str(table.(string), "as")
			ArgsBuilder.Table = util.Cfg.Mysql.TablePrefix+name
			ArgsBuilder.Alias = alias
		}else {
			Ret.Code = err.ERROR_NOT_TABLE
			return
		}
	}
	//where 处理
	if where, ok:= JsonArgs["where"]; ok{
		ArgsBuilder.Where = where
	}
	//字段 处理
	if field, ok:= JsonArgs["field"]; ok{
		ArgsBuilder.Field = field.(string)
	}
	//having 关键字 处理
	if having, ok:= JsonArgs["having"]; ok{
		ArgsBuilder.Having = having.(string)
	}
	//omit 关键字 处理
	if omit, ok:= JsonArgs["omit"]; ok{
		ArgsBuilder.Omit = omit.([]string)
	}
	//distinct 关键字 处理
	if distinct, ok:= JsonArgs["distinct"]; ok{
		ArgsBuilder.Distinct = distinct.([]string)
	}
	//limit 关键字 处理
	if limit, ok:= JsonArgs["limit"]; ok{
		ArgsBuilder.Limit = limit.(int)
	}
	//offset 关键字 处理
	if offset, ok:= JsonArgs["offset"]; ok{
		ArgsBuilder.Offset = offset.(int)
	}
	//order 关键字 处理
	if order, ok:= JsonArgs["order"]; ok{
		ArgsBuilder.Order = order
	}
	//group 关键字 处理
	if group, ok:= JsonArgs["group"]; ok{
		ArgsBuilder.Group = group.([]interface{})
	}
	//join 关键字 处理
	if join, ok:= JsonArgs["join"]; ok{
		ArgsBuilder.Join = join.([]interface{})
	}
	//body 关键字 处理
	if body, ok:= JsonArgs["body"]; ok{
		ArgsBuilder.Body = body
	}
	//id 关键字 处理
	if id, ok:= JsonArgs["id"]; ok{
		ArgsBuilder.Id = id.(int)
	}
}

//初始化 where
func InitWhere(queryBuilder *QueryBuilderS){

	if ArgsBuilder.Where == nil{
		return
	}
	tmp := GeneralQuery(ArgsBuilder.Where, "and")
	queryBuilder.Query = util.MultiTrimSuffix(tmp, []string{"and ", "or "})
}

//初步解析query
func GeneralQuery(mid interface{}, link string)string{
	query := ""

	_,code := util.AssertType(mid)
	if code != util.IsSliceInterface{
		Ret.Code = err.INVALID_PARAMS
		return query
	}
	where := mid.([]interface{})
	lenW := len(where)
	if lenW == 0{
		Ret.Code = err.INVALID_PARAMS
		return query
	}

	for _,val := range where{
		_,code := util.AssertType(val)
		if code != util.IsString{
			query += GeneralQuery(val, link)
			continue
		}
		//判断是否为链接符
		if isExist := util.ContainsString(LinkOption, val.(string)); isExist>=0{
			link =  val.(string)
			tmp , tmpLink := where[1:], ""

			if link == "or"{
				tmpLink = "and "
			}else{
				tmpLink = "or "
			}
			tmpQuery := GeneralQuery(tmp, link)

			tmpQuery = util.MultiTrimSuffix(tmpQuery, []string{link+" ", tmpLink})
			query += "("+tmpQuery+") and "
			break
		}
		if lenW <2{
			Ret.Code = err.INVALID_PARAMS
			break
		}
		query = FurtherQuery(where, lenW)+" "+ link+" "
		break
	}

	return query
}
//进一步解析query
func FurtherQuery(mid []interface{}, lenW int)string{

	field, option, query := mid[0].(string), "=", ""

	var appendVal interface{}
	//当省略操作符 =
	if lenW == 2{
		appendVal, _ = util.AssertType(mid[1])
		queryBuilder.BindValue = append(queryBuilder.BindValue,appendVal )
		query = field + " " + option + " ?"
	}else {
		option = mid[1].(string)
		if isExist := util.ContainsString(WhereOption, option); isExist < 0 {
			Ret.Code = err.INVALID_PARAMS
			return query
		}

		appendVal, code := util.AssertType(mid[2])
		switch option {
		case "in":
			replace := " ( "
			if code != util.IsSliceInterface {
				Ret.Code = err.INVALID_PARAMS
				return query
			}
			tmp := appendVal.([]interface{})
			lenV := len(tmp)
			if lenV == 0 {
				Ret.Code = err.INVALID_PARAMS
				return query
			}

			for key, val := range tmp {
				queryBuilder.BindValue = append(queryBuilder.BindValue, val)
				if key == lenV-1 {
					replace += " ? ) "
				} else {
					replace += "?,"
				}
			}
			query = field + " " + option + replace
		case "field_in_set":
			queryBuilder.BindValue = append(queryBuilder.BindValue, appendVal)
			query = " field_in_set (?, " + field + ") "
		case "between","not between":
			if code != util.IsSliceInterface {
				Ret.Code = err.INVALID_PARAMS
				return query
			}
			tmp := appendVal.([]interface{})
			lenV := len(tmp)
			if lenV < 2  {
				Ret.Code = err.INVALID_PARAMS
				return query
			}
			query = field + " " + option
			for i :=0;i<2 ;i++  {
				queryBuilder.BindValue = append(queryBuilder.BindValue, tmp[i])
				if i ==0{
					query+= " ? and"
				}else {
					query+= " ? "
				}
			}
		default:
			queryBuilder.BindValue = append(queryBuilder.BindValue, appendVal )
			query = field + " " + option + "? "
		}
	}

	return query
}
//初始化字段
func InitField(queryBuilder *QueryBuilderS){
	if ArgsBuilder.Field == "*"{
		return
	}
	field := strings.Split(ArgsBuilder.Field, ",")
	for _,val := range field {
		queryBuilder.Field = append(queryBuilder.Field, strings.TrimSpace(val))
	}
}
//启动排序
func InitOrder(queryBuilder *QueryBuilderS){

	if  ArgsBuilder.Order == nil {
		return
	}
	//对参数进行处理以及类型转换
	var order [][]string
	tmp := ArgsBuilder.Order.([]interface{})
	for  _,val := range tmp {
		var tmpVal []string
		for _, val1 := range val.([]interface{})  {
			tmpVal =  append(tmpVal, val1.(string))
		}
		order =  append(order, tmpVal)
	}

	queryBuilder.Order = map[string][]string{
		"desc" : nil,
		"asc" : nil,
	}
	for _,val := range order{
		if len(val) != 2{
			Ret.Code = err.ERROR_ORDER
			break
		}
		option := val[1]

		if util.ContainsString(OrderOption, option) < 0{
			Ret.Code = err.ERROR_ORDER
			break
		}

		queryBuilder.Order[option] = append(queryBuilder.Order[option], val[0])
	}
}
//初始化字段
func InitJoin(queryBuilder *QueryBuilderS){

	if ArgsBuilder.Join == nil{
		return
	}
	for _,val := range ArgsBuilder.Join  {
		_,code := util.AssertType(val)
		if code != util.IsSliceInterface{
			Ret.Code = err.ERROR_NOT_JOIN
			break
		}
		tmp := val.([]interface{})
		if len(tmp)<3{
			Ret.Code = err.ERROR_NOT_JOIN
			break
		}
		name, alias := util.Split2Str(tmp[1].(string), "as")
		if alias != "" {
			tmp[1] = []string{name, alias}
		}

		queryBuilder.Join = append(queryBuilder.Join, tmp)
	}
	fmt.Println(queryBuilder.Join)
}
