package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//http://xorm.io/ 看看

var db *gorm.DB

func openDB() {
	var err error
	db, err = gorm.Open("mysql", "root:@/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	db.SingularTable(true)
}

type Article struct {
	Id    int
	Title string
	Body  string
}
