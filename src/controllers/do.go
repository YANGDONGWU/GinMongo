package controllers

import (
	"github.com/gin-gonic/gin"
	do "models"
	"net/http"
)

func GetDataList(c *gin.Context){
	book:=do.GetPersonList("1Q84")
	//c.JSON(http.StatusOK, gin.H{
	//	"name": book.Name,
	//	"auther":book.Auther,
	//})
	c.String(http.StatusOK, "Hello World!"+book.Name+book.Auther)
	//c.HTML(http.StatusOK,"list.html",gin.H{
	//	"title":"GIN:用户列表页面",
	//})
}
