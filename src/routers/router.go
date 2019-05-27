package routers

import (
	"github.com/gin-gonic/gin"
	. "controllers"
)

func InitRouter() *gin.Engine{

	router:=gin.Default()

	router.GET("/",GetDataList)

	return router
}
