package driver

import (
	"context"
	"dc/util"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var MongoEngine *mongo.Client

type MongoService struct {
	collection *mongo.Collection
	ctx        context.Context
}


func GetMongoEngine() *mongo.Client {


	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(util.Cfg.Mongodb.MongodbUri))

	err = client.Ping(ctx, readpref.Primary())
	if err !=nil{
		fmt.Println(err.Error())
	}
	return client
}
/**
  链接实例
*/
func (ms MongoService)Connect(db, collection string) MongoService{
	ms.ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	ms.collection = MongoEngine.Database(db).Collection(collection)
	return ms
}


func (ms MongoService)InsertOne( doc interface{}) error {

	_, err := ms.collection.InsertOne(ms.ctx, doc)

	return err
}

//func (ms MongoService)FindOne(obj *models.User, where bson.D) error {
//
//	err := ms.collection.FindOne(ms.ctx, where).Decode(obj)
//
//	fmt.Println(obj)
//	return err
//}
