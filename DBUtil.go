package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type article struct {
	title string
	body  string
}

var db *sql.DB

func initDB() {
	if db != nil {
		return
	}

	db, err := sql.Open("mysql", "root:@/godb")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func insert(a article) {
	//	initDB()

	//	stmt, err := db.Prepare("insert into article(title,body,time,author,cate) values(?,?)")
	//	if err != nil {
	//		log.Fatal(err)
	//		return
	//	}
	//	defer stmt.Close()

	//	stmt.Exec(a.title, a.body)

	fmt.Println(a.title, a.body)
}
