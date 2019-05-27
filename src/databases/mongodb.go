package databases

import (
	"fmt"
	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Book struct{
	Name   string
	Auther string
}

func Find(name string) Book {
	url := "localhost:27017"
	session, err := mgo.Dial(url)
	if err!=nil{
		panic(err)
	}
	defer  session.Close()
	c:=session.DB("allen").C("book")

	result := Book{}
	err = c.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", result.Name)
	fmt.Println("Auther:", result.Auther)

	return result
}