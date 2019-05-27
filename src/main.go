package main

import (
	_ "gopkg.in/mgo.v2"
	router "routers"
)

func main(){

	router:=router.InitRouter()

	router.Run(":8000")
}
