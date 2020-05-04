package driver

import (
	"github.com/streadway/amqp"
	"dc/util"
)


//链接实体
type Connection struct {

	Conn *amqp.Connection
	Channel *amqp.Channel
}

//初始化
func GetRedisEngine() Connection {
	mqService := Connection{}
	mqService.Conn = mqService.dial()
	mqService.Channel = mqService.createChannel()

	return mqService
}


//链接实例
func (conn Connection)dial()*amqp.Connection  {
	ctx, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	util.FailOnError(err, "Failed to connect to RabbitMQ")
	return ctx
}
//打开通道
func  (conn Connection)createChannel()*amqp.Channel{
	channel, err := conn.Conn.Channel()
	util.FailOnError(err, "Failed to open a channel")

	return channel
}

//关闭通道
func (conn Connection)Close(){
	conn.Channel.Close()
	conn.Conn.Close()
}