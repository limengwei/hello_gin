package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func openDB() {
	var err error
	db, err = xorm.NewEngine("mysql", "root:@/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	db.ShowSQL(true)
	db.Logger().SetLevel(core.LOG_DEBUG)
}

type Article struct {
	Id    int
	Title string
	Body  string
}
