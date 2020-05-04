package controllers

import (
	"dc/service/mysql"
	"github.com/gin-gonic/gin"
	"net/http"
)


//返回函数
func Resp(c *gin.Context, resBody mysql.ResBody)  {
	//设置header 因为默认是text/html ,现在设置为application/json
	c.JSON(http.StatusOK, resBody)
}
