package service

import "github.com/go-xorm/xorm"

var (
	Ret ResBody
	tablePrefix = "tg_"
	AffairHandle map[string] *xorm.Session  //事务句柄
)

//返回结构体
type ResBody struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data,omitempty"` //omitempty 当data为null,不返回这个字段
}
//参数接口体
type Args map[string]interface{}

//定义方法接口
//存在接口
type ExistInter interface {
	Exist(arg Args)
}
//数据统计接口
type CountInter interface {
	Count(arg Args)
}
//查询接口
type FetchInter interface {
	Fetch(arg Args)
}
//创建接口
type CreateInter interface {
	Create(arg Args)
}
//修改接口
type UpdateInter interface {
	Update(arg Args)
}
//删除接口
type DeleteInter interface {
	Delete(arg Args)
}
//单条获取接口,可以用于redis
type GetInter interface {
	Get(arg Args)
}
//单条设置
type SetInter interface {
	Set(arg Args)
}
//单条删除
type DelInter interface {
	Del(arg Args)
}
type QueryInter interface {
	Query(arg Args)
}
//定义驱动接口
type HandleInterface interface {
	FetchInter
	CreateInter
	UpdateInter
	DeleteInter
	ExistInter
	//CountInter
	GetInter
	SetInter
	DelInter
	QueryInter

}
type DCService struct {

}

func (sc DCService)Fetch(args string)  {
}

