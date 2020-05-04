package main
//whr-helen 2019

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"strings"
	_ "strings"
)
//用户输入配置
var(
	tables string  //表名称
	path string		//结构体保存路径
	readDir = "../../models"
	filename = "mapper.text"
	tablesArr []string
)

func init() {
	flag.StringVar(&path, "path", "./", "# Structure preservation path")
	flag.StringVar(&tables, "t", "all", "# Table name formats such as - t user, rule, config")
}

func main(){
	flag.Parse()

	if tables != "all"{
		tablesArr = convtables(tables)//转换
	}

	fmt.Println("指定数生成据表:",tables)
	fmt.Println("AutoMapper Struct Start ...")

	strBind, strResolver, strBindFun, strResolverFun:= "//model的bind\n\n", "//model的Resolver\n\n","//model的bind fun\n\n","//model的Resolver 的fun\n\n"

	bindFun := `
func %s(args Args) interface{} {

	model := models.%s{}
	if args != nil{
		mapstructure.Decode(args, &model)
	}
	return model
}
`

	resolverFun := `
func %s(multiModels []interface{}, length int )  {
	data := make([]int, length)
	for k, model:= range multiModels  {
		tmp := (model).(*models.%s)
		data[k] = (*tmp).Id
	}
	Ret.Data = data
}
`
	if len(tablesArr)>0{
		for _, f := range tablesArr {
			fmt.Println(f)
			name := strToUpper(strings.TrimPrefix(f, "tg_"))
			strBind += fmt.Sprintf("tablePrefix+\"%s\":%s,\n",f ,"Bind"+name)
			strResolver += fmt.Sprintf("tablePrefix+\"%s\":%s,\n",f ,"Resolver"+name)
			strBindFun +=  fmt.Sprintf(bindFun,"Bind"+name ,strToUpper(f)) //TgAdminUser
			strResolverFun +=  fmt.Sprintf(resolverFun+"\n\n","Resolver"+name ,strToUpper(f)) //TgAdminUser
		}
	}else{
		files, _ := ioutil.ReadDir(readDir)
		for _, f := range files {
			fmt.Println(f.Name())
			name := strings.TrimPrefix(f.Name(), "tg_")
			name = strings.TrimSuffix(name, ".go")
			upName := strToUpper(name)
			table_name :=strings.TrimSuffix(f.Name(), ".go")
			strBind += fmt.Sprintf("tablePrefix+\"%s\":%s,\n", name ,"Bind"+upName)
			strResolver += fmt.Sprintf("tablePrefix+\"%s\":%s,\n", name ,"Resolver"+upName)
			strBindFun +=  fmt.Sprintf(bindFun,"Bind"+upName ,strToUpper(table_name)) //TgAdminUser
			strResolverFun +=  fmt.Sprintf(resolverFun+"\n\n","Resolver"+upName ,strToUpper(table_name)) //TgAdminUser
		}
	}


	str := strBind+strResolver+strBindFun+strResolverFun

	//创建文件夹
	file,error2 := os.Create(filename)
	if error2 != nil{
		fmt.Println("create file path err:",error2)
	}
	_, err4:=file.Write([]byte(str))
	if err4 != nil{
		fmt.Println("写入文件错误:",err4)
	}

	fmt.Println("End  SUCCESS")
}

//首字母大写
func strToUpper(str string) string {
	var upperStr string
	vv := strings.Split(str, "_")

	for _,v := range vv{
		upperStr += strings.Title(v)
	}
	return upperStr
}

//判断用户输入
func checkpath(path string)bool{
	a := strings.Split(path, "/")
	if len(a)>=2 && a[1] != ""{
		return false
	}else{
		return true
	}

}

//用户输入转换
func convtables(tab string)[]string{
	strArr := strings.Split(tab, ",")
	return strArr
}
