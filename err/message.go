package err

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	ERROR_NOT_TABLE : "表名必须传",
	ERROR_NOT_FUN : "没有此方法",
	ERROR_NOT_DRIVER : "没有此驱动",
	ERROR_NOT_JOIN : "join 参数错误",
	ERROR_ORDER : "order 参数错误",
	ERROR_NOT_WHERE : "查询单条数据,条件必须传",
	ERROR_OPTION : "查询操作符错误",
	ERROR_EXIST : "已存在该数据存在",
	ERROR_NOT_EXIST : "该数据不存在",
	ERROR_NOT_EXIST_QUERY : "该文章不存在",
	ERROR_INSERT_FAIL: "插入数据失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
	ERROR_AUTH_TOKEN : "Token生成失败",
	ERROR_AUTH : "Token错误",
}

func Msg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
