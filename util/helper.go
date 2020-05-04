package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"strings"
)

func Md5Encode(data string) string{
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)

	return  hex.EncodeToString(cipherStr)

}
func MD5Encode(data string) string{
	return strings.ToUpper(Md5Encode(data))
}

func ValidatePasswd(plainpwd, salt, passwd string) bool{
	return Md5Encode(plainpwd+salt)==passwd
}
func MakePasswd(plainpwd, salt string) string{
	return Md5Encode(plainpwd+salt)
}


//错误判断
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//查看字符串是否在切片中
func ContainsString(array []string, val string) (index int){
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}
const (
	IsBoolean uint8 = iota
	IsInteger
	IsUnsigned
	IsFloat
	IsComplex
	IsString
	IsSliceInterface
	IsSliceString

)
//断言返回数据 类型
func AssertType(t interface{})(interface{}, uint8){

	var (
		ret interface{}
		code uint8
	)
	switch t.(type) {
	case string:
		ret,code = t.(string), IsString
		break
	case int:
		ret,code = t.(int), IsInteger
		break
	case json.Number:
		ret,code = t.(json.Number), IsInteger
		break
	case float64:
		ret,code = t.(float64),IsFloat
		break
	case []interface{}:
		ret,code = t.([]interface{}),IsSliceInterface
		break
	case []string:
		ret,code = t.([]string),IsSliceString
		break
	}

	return ret, code
}
//多个去掉字符串后缀
func MultiTrimSuffix(s string, suffix []string) string {
	for _,val := range suffix{
		if strings.HasSuffix(s, val) {
			s = s[:len(s)-len(val)]
		}
	}
	return s
}
//分割字符串,两次分割处理 主要用于 alias
func Split2Str(str, separator string)(string, string){
	name, alias := "",""
	tableArr := strings.Split(str, separator)
	name = strings.TrimSpace(tableArr[0])
	if len(tableArr)>1{
		alias = strings.TrimSpace(tableArr[1])
	}
	return name,alias
}
//把结构体中[]byte convert string
func FormatValue(dataMap *map[string]interface{}) {
	for k, v := range *dataMap {
		//fmt.Printf("key: %s, val: %+v, type: %s\n", k, v, reflect.TypeOf(v))
		switch v.(type) {
		case []byte:
			(*dataMap)[k] = string(v.([]byte))
		case json.RawMessage:
			(*dataMap)[k] = string(v.(json.RawMessage))
		case *uint32:
			(*dataMap)[k] = *v.(*uint32)
		case *int64:
			(*dataMap)[k] = *v.(*int64)
		case *int32:
			(*dataMap)[k] = *v.(*int32)
		case *string:
			(*dataMap)[k] = *v.(*string)
		default:
			(*dataMap)[k] = v
		}
	}
}

