package controllers

import (
	"dc/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spiral/goridge"
	"net"
	"net/http"
	"net/rpc"
)



var userService service.UserService

type User struct {

}
var rp map[string]interface{}
func (u User)Show(c *gin.Context)  {

	resBody := userService.Show(c)


	c.JSON(http.StatusOK, resBody)

}
func (u User)List(c * gin.Context)  {

	fmt.Println(4444)
	conn, err := net.Dial("tcp", ":6001")
	if err != nil {
		panic(err)
	}

	client := rpc.NewClientWithCodec(goridge.NewClientCodec(conn))

	defer client.Close()


	client.Call("App.Hi", "name", &rp)

	fmt.Println(rp)
}
func (u User)Update(c * gin.Context)  {

}
func (u User)Delete(c * gin.Context)  {

}
func (u User)Create(c * gin.Context)  {

}
func (u User)DeleteAll(c * gin.Context)  {

}