package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func openDB() {
	var err error
	session, err = mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
		return
	}
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("godb")

	collecttion := db.C("Users")
	countNum, err := collecttion.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("Things objects count: ", countNum)

	collecttion.Insert(&Person{PHONE: "18811577546", NAME: "zhangzheHero"})
}

type Person struct {
	NAME  string
	PHONE string
}

type Article struct {
	Id    int
	Title string
	Body  string
}
