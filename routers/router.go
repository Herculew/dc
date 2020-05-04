package routers
//路由
import (
	"dc/controllers"
	"dc/util"
	"github.com/gin-gonic/gin"
)

func InitHttpRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(util.Cfg.RunMode)

	//用户表
	userCtr := controllers.User{}

	r.GET("/user", userCtr.Show)
	r.GET("/users", userCtr.List)
	r.POST("/user", userCtr.Create)
	r.PUT("/user/:id", userCtr.Update)
	r.DELETE("/user/:id", userCtr.Delete)
	r.DELETE("/users", userCtr.DeleteAll)


	return r
}
