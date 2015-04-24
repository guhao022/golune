package main

import (
	"fmt"
	//"github.com/astaxie/beego"
	"golune/utils"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	//utils.RegisterDB("localhost", "test")
}

type Person struct {
	Name  string
	Phone string
}

func main() {
	c := utils.NewDB("people")
	err := c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone:", result.Phone)
	//beego.Run()
}
