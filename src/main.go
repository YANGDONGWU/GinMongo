package main

import (
	"fmt"
	_ "gopkg.in/mgo.v2"
	router "routers"
)

func main() {

	router := router.InitRouter()
	fmt.Println("test")
	router.Run(":8000")
}
