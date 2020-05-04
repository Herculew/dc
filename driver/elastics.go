package driver

import (
	"github.com/olivere/elastic/v7"
)

var EsEngine *elastic.Client

//type EsEngine struct {
//	collection *elastic.Client
//}
//
//
//func GetEsEngine() *elastic.Client {
//	client, err := elastic.NewClient()
//
//	if err !=nil{
//		fmt.Println(err.Error())
//	}
//	return client
//}

